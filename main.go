package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/config"
	"backend/controllers"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	config.Conn()

	r.GET("/tanamans", controllers.GetTanaman)
	r.GET("/lahans", controllers.GetLahan)
	r.GET("/inventaris", controllers.GetInventaris)
	r.GET("/schedulers", controllers.GetScheduler)

	r.DELETE("/lahans/:id", controllers.DeleteLahan)
	r.DELETE("/inventaris/:id", controllers.DeleteInventaris)

	r.POST("/schedulers", controllers.TambahScheduler)

	r.PUT("/schedulers/:id/status", controllers.UpdateStatus)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // untuk development lokal
	}

	r.Run(":" + port)
}