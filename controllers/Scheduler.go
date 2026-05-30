package controllers

import (
	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/models"
	"strconv"
)

func GetScheduler(c *gin.Context){
	var scheduler []models.Scheduler
	if err := config.DB.Find(&scheduler).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil data"})
		return
	}
	c.JSON(200, scheduler)
}

func TambahScheduler(c *gin.Context){
	var input models.Scheduler
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Input tidak valid"})
		return
	}
	scheduler := models.Scheduler{
		NamaScheduler: input.NamaScheduler,
		Tanggal: input.Tanggal,
		Status: input.Status,
	}
	if err := config.DB.Create(&scheduler).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal menambahkan data"})
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

func UpdateStatus(c *gin.Context){
	id := c.Param("id")
	var scheduler models.Scheduler

	var body struct {
	Status string `json:"status"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Input tidak valid"})
		return
	}

	if err := config.DB.First(&scheduler, "SchedulerId = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Scheduler tidak ditemukan"})
		return
	}
	scheduler.Status = body.Status
	config.DB.Save(&scheduler)
	c.JSON(200, scheduler)

}