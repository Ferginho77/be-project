package controllers

import (
	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/models"	
)

func GetAktivitas(c *gin.Context) {
	var aktivitas []models.Aktivitas
	if err := config.DB.Find(&aktivitas).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil data"})
		return
	}
	c.JSON(200, aktivitas)
}

func CreateAktivitas(c *gin.Context){
	var aktivitas models.Aktivitas
	if err := c.ShouldBindJSON(&aktivitas); err != nil {
		c.JSON(400, gin.H{"error": "Data tidak valid"})
		return
	}
	if err := config.DB.Create(&aktivitas).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal buat data"})
		return
	}
	c.JSON(201, aktivitas)
}