package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB() {
	dsn := "root:root@tcp(localhost:3306)/product?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func GetDB() *gorm.DB {
	if DB == nil {
		panic("database is not initialized")
	}
	fmt.Println("Database connected successfully!")
	return DB
}