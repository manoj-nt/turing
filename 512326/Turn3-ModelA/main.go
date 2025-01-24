package main

import (
	"log"
	"modela/paramserializer" // Your package's path.

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// We could store tags in a way compatible with GORM explicitly (JSON store).

	// Migration to ensure our schema is up-to-date expressed via structs
	db.AutoMigrate(&paramserializer.User{})

	rawQuery := "user_id=123&name=JohnDoe&age=30" +
		"&address[city]=NewYork&address[state]=NY" +
		"&address[coordinates][lat]=40.7128" +
		"&address[coordinates][lng]=-74.0060" +
		"&tags[]=go&tags[]=backend"

	user, err := paramserializer.SerializeQueryParams(rawQuery)
	if err != nil {
		log.Fatalf("Error serializing: %v", err)
	}

	if err := paramserializer.ValidateUser(user); err != nil {
		log.Fatalf("Validation fail: %v", err)
	}

	// Here the tags field needs special serialization handling
	db.Save(user)

	log.Printf("User: %+v", user)
}
