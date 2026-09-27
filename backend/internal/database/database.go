package database

import (
	"log"

	"backend/internal/models" // Import your new models package
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=myuser password=mypassword dbname=mobileapp port=5432 sslmode=disable TimeZone=UTC"
	
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal("Failed to connect to the database: \n", err)
	}

	log.Println("Database connection successfully established.")

	// --- NEW: Run AutoMigration ---
	log.Println("Running database migrations...")
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database tables: \n", err)
	}
	log.Println("Database migrated successfully.")
}
