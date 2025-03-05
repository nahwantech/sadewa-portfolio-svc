package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

// DB is the global database connection
var DB *gorm.DB

<<<<<<< HEAD
func ConnectDB() {

	// Explicitly load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}


	log.Println("DB HOST : ", os.Getenv("DB_HOST"))
	log.Println("DB PORT : ", os.Getenv("DB_PORT"))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
=======
// InitDB initializes the database connection
func InitDB() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	
	// Read environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
>>>>>>> feature/db-integration/init

	// Validate required fields
	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		log.Fatal("Database environment variables are not set properly")
	}

<<<<<<< HEAD
	// Configure connection pool 
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)
=======
	// Create DSN string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)
>>>>>>> feature/db-integration/init

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[ERROR] Failed to connect to database: %v", err)
	}

	log.Println("[INFO] Database connected successfully")
}
