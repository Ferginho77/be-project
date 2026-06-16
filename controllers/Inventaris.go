package controllers

import (
	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/models"
	"strconv"
		
)

func GetInventaris(c *gin.Context) {
	var inventaris []models.Inventaris
	if err := config.DB.Find(&inventaris).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil data"})
		return
	}
	c.JSON(200, inventaris)
}

func DeleteInventaris(c *gin.Context) {
	var inventaris models.Inventaris

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := config.DB.Where("InventarisId = ?", id).Take(&inventaris).Error; err != nil {
		c.JSON(404, gin.H{"error": "Inventaris tidak ditemukan"})
		return
	}

	if err := config.DB.Delete(&inventaris).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal hapus data"})
		return
	}

	c.JSON(200, gin.H{"message": "Inventaris berhasil dihapus"})
}

func CreateInventaris(c *gin.Context) {
	var inventaris models.Inventaris
	if err := c.ShouldBindJSON(&inventaris); err != nil {
		c.JSON(400, gin.H{"error": "Data tidak valid"})
		return
	}
	if err := config.DB.Create(&inventaris).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal buat data"})
		return
	}
	c.JSON(201, inventaris)
}


func UpdateInventaris(c *gin.Context) {
	var inventaris models.Inventaris
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID tidak valid"})
		return
	}
	if err := config.DB.Where("InventarisId = ?", id).Take(&inventaris).Error; err != nil {
		c.JSON(404, gin.H{"error": "Inventaris tidak ditemukan"})
		return
	}
	var input models.Inventaris
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Data tidak valid"})
		return
	}
	inventaris.NamaBarang = input.NamaBarang
	inventaris.Jenis = input.Jenis
	inventaris.Stok = input.Stok
	if err := config.DB.Save(&inventaris).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal update data"})
		return
	}
	c.JSON(200, inventaris)
}