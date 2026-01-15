// Package utils provides utility functions for various common tasks.
package utils

import (
	"fmt"
	"os"

	m "sentra-medika/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		err = godotenv.Load("../.env")
	}

	if err != nil {
		panic("❌ An error occurred while loading .env file, please check if file exists in root or parent directory.")
	}

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		panic("❌ Database URL is not set in the environment variables.")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("❌ Failed to connect to the database.")
	}

	fmt.Println("✅ Successfully connected to the database.")

	err = database.AutoMigrate(&m.Users{}, &m.MedicalRecords{})

	if err != nil {
		panic("❌ Failed to run database migrations.")
	}

	DB = database
	return database
}