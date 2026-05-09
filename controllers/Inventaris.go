package controllers

import (
	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/models"
		
)

func GetInventaris(c *gin.Context) {
	var inventaris []models.Inventaris
	if err := config.DB.Find(&inventaris).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil data"})
		return
	}
	c.JSON(200, inventaris)
}