package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=91.108.104.69 user=postgres password=yourpassword dbname=gqlgendb port=5432 sslmode=disable"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.info), // Enable logging for queries
	})

	if err != null {
		log.Fatal("Failed to migrate database : ", err)
	}

	fmt.Println("Connected to PostgreSQL database")
}