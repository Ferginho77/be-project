package controllers

import (
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