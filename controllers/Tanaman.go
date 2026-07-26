package controllers

import (
	"rest-api/config"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func GetTanaman(c *gin.Context) {
	var tanaman []models.Tanaman
	if err := config.DB.Find(&tanaman).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil data"})
		return
	}
	c.JSON(200, tanaman)
}

func DeleteTanaman(c *gin.Context) {
	var tanaman models.Tanaman
	id := c.Param("id")
	if err := config.DB.Where("TanamanId = ?", id).First(&tanaman).Error; err != nil {
		c.JSON(404, gin.H{"error": "Tanaman tidak ditemukan"})
		return
	}
	if err := config.DB.Delete(&tanaman).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal hapus data"})
		return
	}
	c.JSON(200, gin.H{"message": "Tanaman berhasil dihapus"})
}

func CreateTanaman(c *gin.Context) {
	var tanaman models.Tanaman
	if err := c.ShouldBindJSON(&tanaman); err != nil {
		c.JSON(400, gin.H{"error": "Gagal parsing data"})
		return
	}
	if err := config.DB.Create(&tanaman).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal input data"})
		return
	}
	c.JSON(200, gin.H{"message": "Tanaman berhasil ditambahkan"})
}

func UpdateTanaman(c *gin.Context) {
	var tanaman models.Tanaman
	id := c.Param("id")
	if err := config.DB.Where("TanamanId = ?", id).First(&tanaman).Error; err != nil {
		c.JSON(404, gin.H{"error": "Tanaman tidak ditemukan"})
		return
	}
	if err := c.ShouldBindJSON(&tanaman); err != nil {
		c.JSON(400, gin.H{"error": "Gagal parsing data"})
		return
	}
	if err := config.DB.Save(&tanaman).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal update data"})
		return
	}
	c.JSON(200, gin.H{"message": "Tanaman berhasil diupdate"})
}
