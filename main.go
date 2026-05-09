package main

import (
    "rest-api/config"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "rest-api/controllers"
)

func main() {
    r := gin.Default()

    // Setup CORS agar bisa diakses oleh React (biasanya port 3000 atau 5173)
    r.Use(cors.Default())

    // Koneksi Database
    config.Conn()

    // Route Get Semua Data
    r.GET("/tanamans", controllers.GetTanaman)
    r.GET("/lahans", controllers.GetLahan)
    r.GET("/inventaris", controllers.GetInventaris)

    // Route Delete
    r.DELETE("/lahans/:id", controllers.DeleteLahan)

    r.Run(":8080") // Default di port 8080
}