package controllers

import (
	"fmt"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/models"
	"strconv"
)

func GetPenanaman(c *gin.Context) {
	var penanaman []models.Penanaman
	if err := config.DB.Find(&penanaman).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil data"})
		return
	}
	c.JSON(200, penanaman)
}

func DeletePenanaman(c *gin.Context) {
	var penanaman models.Penanaman
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := config.DB.Where("PenanamanId = ?", id).Take(&penanaman).Error; err != nil {
		c.JSON(404, gin.H{"error": "Penanaman tidak ditemukan"})
		return
	}
	if err := config.DB.Delete(&penanaman).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal hapus data"})
		return
	}
	c.JSON(200, gin.H{"message": "Penanaman berhasil dihapus"})
}

// scheduleTemplate mendefinisikan proporsi hari dan nama task berdasarkan siklus 30 hari.
// Proporsi akan discale otomatis terhadap UmurPanen tanaman sesungguhnya.
type scheduleTemplate struct {
	DayRatio float64 // rasio terhadap umur panen (0.0 – 1.0)
	Nama     string
}

var templateTasks = []scheduleTemplate{
	{DayRatio: 1.0 / 30.0, Nama: "Penyiraman Awal"},
	{DayRatio: 3.0 / 30.0, Nama: "Pemupukan Dasar"},
	{DayRatio: 7.0 / 30.0, Nama: "Pengecekan Hama"},
	{DayRatio: 10.0 / 30.0, Nama: "Pruning / Pemangkasan"},
	{DayRatio: 15.0 / 30.0, Nama: "Pemupukan Susulan"},
	{DayRatio: 20.0 / 30.0, Nama: "Monitoring Buah"},
	{DayRatio: 25.0 / 30.0, Nama: "Pengecekan Kematangan"},
	{DayRatio: 29.0 / 30.0, Nama: "Persiapan Panen"},
}

func CreatePenanaman(c *gin.Context) {
	var penanaman models.Penanaman
	if err := c.ShouldBindJSON(&penanaman); err != nil {
		c.JSON(400, gin.H{"error": "Data tidak valid"})
		return
	}

	// Ambil data tanaman untuk mendapatkan UmurPanen
	var tanaman models.Tanaman
	if err := config.DB.Where("TanamanId = ?", penanaman.TanamanId).First(&tanaman).Error; err != nil {
		c.JSON(404, gin.H{"error": "Tanaman tidak ditemukan"})
		return
	}

	// Parse TanggalTanam
	tanggalTanam, err := time.Parse("2006-01-02", penanaman.TanggalTanam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Format TanggalTanam tidak valid (gunakan YYYY-MM-DD)"})
		return
	}

	umurPanen := int(tanaman.UmurPanen)
	if umurPanen <= 0 {
		umurPanen = 30 // default 30 hari jika tidak diset
	}

	// Mulai transaksi DB — jika ada yang gagal, semua dibatalkan
	tx := config.DB.Begin()
	if tx.Error != nil {
		c.JSON(500, gin.H{"error": "Gagal memulai transaksi"})
		return
	}

	// 1. Simpan penanaman
	if err := tx.Create(&penanaman).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Gagal simpan data penanaman"})
		return
	}

	// 2. Generate schedulers berdasarkan template × UmurPanen
	for _, tmpl := range templateTasks {
		dayOffset := int(math.Round(tmpl.DayRatio * float64(umurPanen)))
		if dayOffset < 1 {
			dayOffset = 1
		}
		jadwalTanggal := tanggalTanam.AddDate(0, 0, dayOffset)

		scheduler := models.Scheduler{
			NamaScheduler: fmt.Sprintf("Hari %d - %s", dayOffset, tmpl.Nama),
			Tanggal:       jadwalTanggal.Format("2006-01-02"),
			Status:        "Pending",
			PenanamanId:   penanaman.PenanamanId,
		}

		if err := tx.Create(&scheduler).Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": fmt.Sprintf("Gagal membuat scheduler '%s', penanaman dibatalkan", tmpl.Nama)})
			return
		}
	}

	// 3. Auto-update StatusLahan menjadi "Aktif"
	if penanaman.LahanId != 0 {
		if err := tx.Model(&models.Lahan{}).
			Where("LahanId = ?", penanaman.LahanId).
			Update("StatusLahan", "Aktif").Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Gagal update status lahan, penanaman dibatalkan"})
			return
		}
	}

	// Commit transaksi
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Gagal commit transaksi"})
		return
	}

	c.JSON(201, gin.H{
		"penanaman":       penanaman,
		"schedulerDibuat": len(templateTasks),
		"umurPanen":       umurPanen,
		"message":         fmt.Sprintf("Penanaman berhasil dibuat dengan %d jadwal otomatis", len(templateTasks)),
	})
}

func UpdatePenanaman(c *gin.Context) {
	var penanaman models.Penanaman
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := config.DB.Where("PenanamanId = ?", id).Take(&penanaman).Error; err != nil {
		c.JSON(404, gin.H{"error": "Penanaman tidak ditemukan"})
		return
	}
	var input models.Penanaman
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Data tidak valid"})
		return
	}

	tx := config.DB.Begin()
	if tx.Error != nil {
		c.JSON(500, gin.H{"error": "Gagal memulai transaksi database"})
		return
	}

	penanaman.TanamanId = input.TanamanId
	penanaman.TanggalTanam = input.TanggalTanam
	penanaman.RencanaPanen = input.RencanaPanen
	penanaman.JumlahBibit = input.JumlahBibit
	penanaman.Fase = input.Fase
	penanaman.Status = input.Status

	if err := tx.Save(&penanaman).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Gagal update data penanaman"})
		return
	}

	// Update status lahan dan scheduler yang terkait secara otomatis dan konsisten
	if penanaman.Status == "Selesai" || penanaman.Status == "Gagal" {
		if penanaman.LahanId != 0 {
			if err := tx.Model(&models.Lahan{}).
				Where("LahanId = ?", penanaman.LahanId).
				Update("StatusLahan", "Kosong").Error; err != nil {
				tx.Rollback()
				c.JSON(500, gin.H{"error": "Gagal update status lahan"})
				return
			}
		}
		if err := tx.Model(&models.Scheduler{}).
			Where("PenanamanId = ? AND Status = ?", penanaman.PenanamanId, "Pending").
			Update("Status", "Dibatalkan").Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Gagal update status scheduler"})
			return
		}
	} else if penanaman.Status == "Aktif" || penanaman.Status == "Panen" {
		if penanaman.LahanId != 0 {
			if err := tx.Model(&models.Lahan{}).
				Where("LahanId = ?", penanaman.LahanId).
				Update("StatusLahan", "Aktif").Error; err != nil {
				tx.Rollback()
				c.JSON(500, gin.H{"error": "Gagal update status lahan"})
				return
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Gagal commit transaksi"})
		return
	}

	c.JSON(200, penanaman)
}