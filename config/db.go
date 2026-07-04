// CONFIG ONLINE

package config

import (
"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "xu1nvg.h.filess.io"
	port     = "3307"
	user     = "goland_settingeat"
	password = "edc5b3d3d8bfc6bc05b7a052bc24905752e21aa1"
	dbname   = "goland_settingeat"
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

	DB = db

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