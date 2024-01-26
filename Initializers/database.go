package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global var
var DB *gorm.DB

func DatabaseConnection() {
	var err error
	dsn := os.Getenv("DATABASE_STRING")
	// CONNECTING TO DB
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the DB")
	}
}
