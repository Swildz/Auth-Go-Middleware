package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_gojwtsiddiq"))
	if err != nil {
		fmt.Printf("Gagal koneksi databases")
	}

	db.AutoMigrate(&User{})

	DB = db

}
