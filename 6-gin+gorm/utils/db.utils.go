package utils

import (
	"github.com/diegofonseca/go-way/models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func SetupDB() (*gorm.DB, error) {

	dsn := "go:go@tcp(mysql:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.Car{})

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
}
