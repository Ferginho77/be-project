// CONFIG ONLINE

package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	SQLDB *sql.DB
)

const (
	host     = "1aseyi.h.filess.io"
	port     = "3307"
	user     = "tumbura_lifedream"
	password = "35b858ae13a183abea8c340928a8f07c2bc63d8b"
	dbname   = "tumbura_lifedream"
)

func Conn() {

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
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Setting Connection Pool
	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	DB = db
	SQLDB = sqlDB

	fmt.Println("Database Connected")
}

// CONFIG LOCAL

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