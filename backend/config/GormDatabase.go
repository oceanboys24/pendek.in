package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pendek.in/models"
)

var DbClient *gorm.DB

func ConnectDatabaseNeon()  {
	dsn := os.Getenv("DATABASE_URL")
	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot Connect Database", err)
	}

	db.AutoMigrate(&models.ShortModel{})

	DbClient = db
}