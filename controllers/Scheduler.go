package controllers

import (
	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/models"
	"strconv"
	"fmt"

)

func GetScheduler(c *gin.Context){
	var scheduler []models.Scheduler
	if err := config.DB.Find(&scheduler).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil data"})
		return
	}
	c.JSON(200, scheduler)
}

func CreateScheduler(c *gin.Context) {
	var scheduler models.Scheduler
	
	if err := c.ShouldBindJSON(&scheduler); err != nil {
		c.JSON(400, gin.H{"error": "Data tidak valid"})
		return
	}
	if err := config.DB.Create(&scheduler).Error; err != nil {
		c.JSON(500, gin.H{"error": "Gagal buat data"})
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

func UpdateStatus(c *gin.Context) {
	id := c.Param("id")

	var scheduler models.Scheduler

	var body struct {
		Status string `json:"Status"`
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

	config.DB.Model(&models.Scheduler{}).
    Where("SchedulerId = ?", id).
    Update("Status", body.Status)

	c.JSON(200, scheduler)
}

func UpdateScheduler(c *gin.Context) {
	id := c.Param("id")

	var scheduler models.Scheduler

	if err := c.BindJSON(&scheduler); err != nil {
		c.JSON(400, gin.H{
			"error": "Input tidak valid",
		})
		return
	}

	fmt.Printf("DATA DARI FRONTEND: %+v\n", scheduler)

	if err := config.DB.
		Model(&models.Scheduler{}).
		Where("SchedulerId = ?", id).
		Updates(map[string]interface{}{
			"NamaScheduler": scheduler.NamaScheduler,
			"Tanggal":       scheduler.Tanggal,
			"Status":        scheduler.Status,
		}).Error; err != nil {

		fmt.Println("ERROR DATABASE:", err)

		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.First(&scheduler, "SchedulerId = ?", id)

	c.JSON(200, scheduler)
}