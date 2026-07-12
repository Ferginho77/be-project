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

	if err := config.DB.Create(&produksi).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal membuat data produksi"})
		return
	}

	if produksi.PenanamanId != 0 {
		if err := config.DB.Model(&models.Penanaman{}).
			Where("PenanamanId = ?", produksi.PenanamanId).
			Update("Status", "Selesai").Error; err != nil {
			config.DB.Rollback()
			c.JSON(500, gin.H{"error": "Gagal update status Penanaman, penanaman dibatalkan"})
			return
		}
	}


	c.JSON(200, gin.H{"message": "Data produksi berhasil ditambahkan"})
}
