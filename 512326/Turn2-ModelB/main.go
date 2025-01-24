package main

import (
	"log"
	"modela/paramserializer"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize GORM with a Sqlite database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto-migrate User struct to create the users table
	db.AutoMigrate(&paramserializer.User{})

	// Example query string
	rawQuery := "user_id=123&name=JohnDoe&age=30"

	// Serialize query parameters to User struct
	user, err := paramserializer.SerializeQueryParams(rawQuery)
	if err != nil {
		log.Fatalf("failed to serialize query params: %v", err)
	}

	// Create or update the user record in the database
	db.Save(user)
}
