package controllers

import (
	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/models"
	"strconv"
		
)

func GetLahan(c *gin.Context) {
	var lahan []models.Lahan
	if err := config.DB.Find(&lahan).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil data"})
		return
	}
	c.JSON(200, lahan)
}

func DeleteLahan(c *gin.Context) {
	var lahan models.Lahan

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := config.DB.Where("LahanId = ?", id).Take(&lahan).Error; err != nil {
		c.JSON(404, gin.H{"error": "Lahan tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&lahan).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal hapus data"})
		return
	}

	c.JSON(200, gin.H{"message": "Lahan berhasil dihapus"})
}