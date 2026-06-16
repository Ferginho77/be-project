package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rest-api/controllers"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/tanamans", controllers.GetTanaman)
	r.GET("/lahans", controllers.GetLahan)
	r.GET("/inventaris", controllers.GetInventaris)
	r.GET("/schedulers", controllers.GetScheduler)
	r.GET("/penanamans", controllers.GetPenanaman)

	r.DELETE("/lahans/:id", controllers.DeleteLahan)
	r.DELETE("/inventaris/:id", controllers.DeleteInventaris)
	r.DELETE("/penanamans/:id", controllers.DeletePenanaman)
	r.DELETE("/schedulers/:id", controllers.DeleteScheduler)

	r.POST("/schedulers", controllers.CreateScheduler)
	r.POST("/inventaris", controllers.CreateInventaris)
	r.POST("/schedulers/:id", controllers.UpdateStatus)


	r.PUT("/schedulers/:id/update", controllers.UpdateScheduler)
	r.PUT("/inventaris/:id/update", controllers.UpdateInventaris)


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // untuk development lokal
	}

	r.Run(":" + port)
}