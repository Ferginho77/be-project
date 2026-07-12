package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	SQLDB *sql.DB
)

func Conn() {

	// Load .env hanya jika ada (development lokal)
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Validasi agar mudah mendeteksi kesalahan konfigurasi
	if host == "" || port == "" || user == "" || dbname == "" {
		log.Fatal("Environment Variable database belum lengkap.")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Connection Pool
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	DB = db
	SQLDB = sqlDB

	fmt.Println("✅ Database Connected")
}


// package config

// import (
// 	"fmt"
// 	"log"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func Conn() {
// 	dsn := "root:@tcp(127.0.0.1:3306)/goland?parseTime=true"
// 	var err error

// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal("Failed to connect to database: ", err)
// 	} else {
// 		fmt.Println("Database connection successful")
// 	}
// }