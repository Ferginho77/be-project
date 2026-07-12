package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rest-api/controllers"
	"rest-api/config"
)

func main() {
	r := gin.Default()

	// 1. Inisialisasi Database
	config.Conn()

	// 2. Middleware
	r.Use(cors.Default())

	// 3. Routing (Dikelompokkan berdasarkan entitas/resource)
	
	// --- Tanaman ---
	r.GET("/tanamans", controllers.GetTanaman)

	// LOGIN
	r.POST("/login", controllers.Login)

	// --- Lahan ---
	lahan := r.Group("/lahans")
	{
		lahan.GET("", controllers.GetLahan)
		lahan.POST("", controllers.CreateLahan)
		lahan.PUT("/:id/update", controllers.UpdateLahan)
		lahan.DELETE("/:id", controllers.DeleteLahan)
	}

	// --- Inventaris ---
	inventaris := r.Group("/inventaris")
	{
		inventaris.GET("", controllers.GetInventaris)
		inventaris.POST("", controllers.CreateInventaris)
		inventaris.PUT("/:id/update", controllers.UpdateInventaris)
		inventaris.DELETE("/:id", controllers.DeleteInventaris)
	}

	// --- Scheduler ---
	scheduler := r.Group("/schedulers")
	{
		scheduler.GET("", controllers.GetScheduler)
		scheduler.POST("", controllers.CreateScheduler)
		scheduler.POST("/:id", controllers.UpdateStatus) 
		scheduler.PUT("/:id/update", controllers.UpdateScheduler)
		scheduler.DELETE("/:id", controllers.DeleteScheduler)
	}

	// --- Penanaman ---
	penanaman := r.Group("/penanamans")
	{
		penanaman.GET("", controllers.GetPenanaman)
		penanaman.POST("", controllers.CreatePenanaman)
		penanaman.PUT("/:id/update", controllers.UpdatePenanaman)
		penanaman.DELETE("/:id", controllers.DeletePenanaman)
	}

	// --- Produksi ---
	produksi := r.Group("/produksi")
	{
		produksi.GET("", controllers.GetProduksi)
		produksi.POST("", controllers.CreateProduksi)
	}

	// --- Aktivitas ---
	aktivitas := r.Group("/aktivitas")
	{
		aktivitas.GET("", controllers.GetAktivitas)
		aktivitas.POST("", controllers.CreateAktivitas)
	}

	// 4. Setup Port & Jalankan Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default untuk development lokal
	}

	r.Run(":" + port)
}