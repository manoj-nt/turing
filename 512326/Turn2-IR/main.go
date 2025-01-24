package main

import (
	"log"
	"modela/paramserializer" // Your package's path.

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize GORM with a SQLite database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto-migrate the User struct, which includes the embedded Address struct
	db.AutoMigrate(&paramserializer.User{})

	// Example query string with nested parameters
	rawQuery := "user_id=123&name=JohnDoe&age=30&address[city]=NewYork&address[state]=NY"

	// Serialize query parameters to User struct
	user, err := paramserializer.SerializeQueryParams(rawQuery)
	if err != nil {
		log.Fatalf("failed to serialize query params: %v", err)
	}

	// Validate the parsed User struct
	if err := paramserializer.ValidateUser(user); err != nil {
		log.Fatalf("validation failed: %v", err)
	}

	// Create or update the user record in the database
	db.Save(user)

	// Print the result for confirmation
	log.Printf("User: %+v", user)
}
