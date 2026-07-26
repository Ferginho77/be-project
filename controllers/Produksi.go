package controllers

import (
	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/models"
	
		
)

func GetProduksi(c *gin.Context){
	var produksi []models.Produksi
	if err := config.DB.Find(&produksi).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil data"})
		return
	}
	c.JSON(200, produksi)
	
}

func CreateProduksi(c *gin.Context) {
	var produksi models.Produksi

	if err := c.ShouldBindJSON(&produksi); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Mulai Transaksi Database
	tx := config.DB.Begin()
	if tx.Error != nil {
		c.JSON(500, gin.H{"error": "Gagal memulai transaksi database"})
		return
	}

	// A. Insert data ke tabel Produksi
	if err := tx.Create(&produksi).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Gagal membuat data produksi"})
		return
	}

	if produksi.PenanamanId != 0 {
		var penanaman models.Penanaman
		if err := tx.First(&penanaman, "PenanamanId = ?", produksi.PenanamanId).Error; err != nil {
			tx.Rollback()
			c.JSON(404, gin.H{"error": "Penanaman tidak ditemukan"})
			return
		}

		// B. Update Penanaman: Status = Selesai
		if err := tx.Model(&models.Penanaman{}).
			Where("PenanamanId = ?", produksi.PenanamanId).
			Update("Status", "Selesai").Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Gagal update status Penanaman"})
			return
		}

		// C. Update Scheduler: Semua scheduler yang masih Pending pada Penanaman tersebut berubah menjadi Archived/Cancelled (Dibatalkan)
		if err := tx.Model(&models.Scheduler{}).
			Where("PenanamanId = ? AND Status = ?", produksi.PenanamanId, "Pending").
			Update("Status", "Dibatalkan").Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Gagal update status scheduler"})
			return
		}

		// D. Update Lahan: StatusLahan = Kosong
		if penanaman.LahanId != 0 {
			if err := tx.Model(&models.Lahan{}).
				Where("LahanId = ?", penanaman.LahanId).
				Update("StatusLahan", "Kosong").Error; err != nil {
				tx.Rollback()
				c.JSON(500, gin.H{"error": "Gagal update status lahan"})
				return
			}
		}
	}

	// Commit Transaksi
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Gagal commit transaksi"})
		return
	}

	c.JSON(200, gin.H{"message": "Data produksi berhasil disimpan dan status lahan telah diperbarui"})
}
