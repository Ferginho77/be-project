package main

import (
	"os"
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"rest-api/controllers"
	"rest-api/config"
	"rest-api/agents"
	"rest-api/middleware"
)

func main() {
	r := gin.Default()
	 if err := godotenv.Load(); err != nil {
        log.Println("Warning: file .env tidak ditemukan")
    }

	// 1. Inisialisasi Database
	config.Conn()

	// 2. Middleware
	r.Use(cors.Default())

	// 3. Routing (Dikelompokkan berdasarkan entitas/resource)

	agent := agents.NewMakersAgent()

	// Inisialisasi controller dengan menginjeksikan agent
	aiController := controllers.NewAIController(agent)
	r.POST("/api/v1/analyze", aiController.Analyze)
	
	// --- Tanaman ---
	tanaman := r.Group("/tanamans")
	tanaman.Use(middleware.AuthMiddleware())
	{
		tanaman.GET("", controllers.GetTanaman)
		tanaman.POST("", controllers.CreateTanaman)
		tanaman.PUT("/:id/update", controllers.UpdateTanaman)
		tanaman.DELETE("/:id", controllers.DeleteTanaman)
	}

	// LOGIN & REGISTER
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)


	// --- Lahan ---
	lahan := r.Group("/lahans")
	lahan.Use(middleware.AuthMiddleware())
	{
		lahan.GET("", controllers.GetLahan)
		lahan.POST("", controllers.CreateLahan)
		lahan.GET("/:id/control", controllers.GetLahanControl)
		lahan.PUT("/:id/update", controllers.UpdateLahan)
		lahan.DELETE("/:id", controllers.DeleteLahan)
	}

	// --- Inventaris ---
	inventaris := r.Group("/inventaris")
	inventaris.Use(middleware.AuthMiddleware())
	{
		inventaris.GET("", controllers.GetInventaris)
		inventaris.POST("", controllers.CreateInventaris)
		inventaris.PUT("/:id/update", controllers.UpdateInventaris)
		inventaris.DELETE("/:id", controllers.DeleteInventaris)
	}

	// --- Scheduler ---
	scheduler := r.Group("/schedulers")
	scheduler.Use(middleware.AuthMiddleware())
	{
		scheduler.GET("", controllers.GetScheduler)
		scheduler.POST("", controllers.CreateScheduler)
		scheduler.POST("/:id", controllers.UpdateStatus) 
		scheduler.PUT("/:id/update", controllers.UpdateScheduler)
		scheduler.DELETE("/:id", controllers.DeleteScheduler)
	}

	// --- Penanaman ---
	penanaman := r.Group("/penanamans")
	penanaman.Use(middleware.AuthMiddleware())
	{
		penanaman.GET("", controllers.GetPenanaman)
		penanaman.POST("", controllers.CreatePenanaman)
		penanaman.PUT("/:id/update", controllers.UpdatePenanaman)
		penanaman.DELETE("/:id", controllers.DeletePenanaman)
	}

	// --- Produksi ---
	produksi := r.Group("/produksi")
	produksi.Use(middleware.AuthMiddleware())
	{
		produksi.GET("", controllers.GetProduksi)
		produksi.POST("", controllers.CreateProduksi)
	}

	// --- Aktivitas ---
	aktivitas := r.Group("/aktivitas")
	aktivitas.Use(middleware.AuthMiddleware())
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