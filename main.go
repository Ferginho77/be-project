package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"rest-api/config"
	"rest-api/controllers"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	config.Conn()

	r.GET("/tanamans", controllers.GetTanaman)
	r.GET("/lahans", controllers.GetLahan)
	r.GET("/inventaris", controllers.GetInventaris)
	r.GET("/schedulers", controllers.GetScheduler)
	r.GET("/penanamans", controllers.GetPenanaman)

	r.DELETE("/lahans/:id", controllers.DeleteLahan)
	r.DELETE("/inventaris/:id", controllers.DeleteInventaris)
	r.DELETE("/penanamans/:id", controllers.DeletePenanaman)

	r.POST("/schedulers", controllers.CreateScheduler)
	r.POST("/inventaris", controllers.CreateInventaris)

	r.PUT("/schedulers/:id/status", controllers.UpdateStatus)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // untuk development lokal
	}

	r.Run(":" + port)
}