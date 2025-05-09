package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabaseNeon()  {
	dsn := os.Getenv("DATABASE_URL")
	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot Connect Database", err)
	}

	DB = db
}