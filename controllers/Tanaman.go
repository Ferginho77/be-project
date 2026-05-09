package controllers

import (
	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/models"
)

func GetTanaman(c *gin.Context) {
	var tanaman []models.Tanaman
	if err := config.DB.Find(&tanaman).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil data"})
		return
	}
	c.JSON(200, tanaman)
}

func deleteTanaman(c *gin.Context) {
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