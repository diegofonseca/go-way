package utils

import (
	"github.com/diegofonseca/go-way/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func SetupDB() (*gorm.DB, error) {

	dsn := os.Getenv("DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Car{})

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
}
