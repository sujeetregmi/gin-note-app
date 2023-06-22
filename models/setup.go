package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=regmi password=regmi123 dbname=notes port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database.")
	}
	DB = database
}

func DbMigrate() {
	DB.AutoMigrate(&Note{})
	DB.AutoMigrate(&User{})
}
