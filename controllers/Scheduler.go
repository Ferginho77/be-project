package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/models"
)

func GetScheduler(c *gin.Context){
	var scheduler []models.Scheduler
	if err := config.DB.Find(&scheduler).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil data"})
		return
	}
	c.JSON(200, scheduler)
}

func CreateScheduler(c *gin.Context) {
	var scheduler models.Scheduler
	
	if err := c.ShouldBindJSON(&scheduler); err != nil {
		c.JSON(400, gin.H{"error": "Data tidak valid"})
		return
	}
	if err := config.DB.Create(&scheduler).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal buat data"})
		return
	}
	c.JSON(201, scheduler)
}

func DeleteScheduler(c *gin.Context) {
	var scheduler models.Scheduler
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := config.DB.Where("SchedulerId = ?", id).Take(&scheduler).Error; err != nil {
		c.JSON(404, gin.H{"error": "Scheduler tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&scheduler).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal hapus data"})
		return
	}
	c.JSON(200, gin.H{"message": "Scheduler berhasil dihapus"})
}

// jenisAktivitasFromScheduler memetakan nama agenda ke enum JenisAktivitas yang valid.
func jenisAktivitasFromScheduler(namaScheduler string) string {
	// Enum yang valid di database: 'Pemupukan', 'Penyiraman', 'Pengobatan'
	lowered := strings.ToLower(namaScheduler)
	keywords := map[string]string{
		"pupuk":    "Pemupukan",
		"siram":    "Penyiraman",
		"air":      "Penyiraman",
		"obat":     "Pengobatan",
		"semprot":  "Pengobatan",
		"hama":     "Pengobatan",
		"penyakit": "Pengobatan",
	}
	for keyword, jenis := range keywords {
		if strings.Contains(lowered, keyword) {
			return jenis
		}
	}
	// Default fallback
	return "Pemupukan"
}

func UpdateStatus(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Status string `json:"Status"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Input tidak valid"})
		return
	}

	var scheduler models.Scheduler
	if err := config.DB.First(&scheduler, "SchedulerId = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Scheduler tidak ditemukan"})
		return
	}

	// Jalankan dalam satu transaksi atomik
	tx := config.DB.Begin()
	if tx.Error != nil {
		c.JSON(500, gin.H{"error": "Gagal memulai transaksi"})
		return
	}

	// 1. Update status scheduler
	if err := tx.Model(&models.Scheduler{}).
		Where("SchedulerId = ?", id).
		Update("Status", body.Status).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Gagal update status"})
		return
	}

	// 2. Jika status menjadi "Done" (sesuai ENUM database), buat record aktivitas baru
	if body.Status == "Done" {
		tanggalSelesai := time.Now().Format("2006-01-02")
		aktivitas := models.Aktivitas{
			JenisAktivitas: jenisAktivitasFromScheduler(scheduler.NamaScheduler),
			Tanggal:        tanggalSelesai,
			Keterangan:     fmt.Sprintf("Agenda '%s' selesai (Blok #%d)", scheduler.NamaScheduler, scheduler.PenanamanId),
			PenanamanId:    scheduler.PenanamanId,
			SchedulerId:    scheduler.SchedulerId, // referensi ke agenda asal
		}
		if err := tx.Create(&aktivitas).Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Gagal mencatat aktivitas"})
			return
		}
	}

	tx.Commit()

	// Reload scheduler yang sudah diupdate
	config.DB.First(&scheduler, "SchedulerId = ?", id)
	c.JSON(200, scheduler)
}

func UpdateScheduler(c *gin.Context) {
	id := c.Param("id")

	var scheduler models.Scheduler

	if err := c.BindJSON(&scheduler); err != nil {
		c.JSON(400, gin.H{
			"error": "Input tidak valid",
		})
		return
	}

	if err := config.DB.
		Model(&models.Scheduler{}).
		Where("SchedulerId = ?", id).
		Updates(map[string]interface{}{
			"NamaScheduler": scheduler.NamaScheduler,
			"Tanggal":       scheduler.Tanggal,
			"Status":        scheduler.Status,
		}).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.First(&scheduler, "SchedulerId = ?", id)

	c.JSON(200, scheduler)
}