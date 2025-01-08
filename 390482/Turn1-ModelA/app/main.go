// main.go
package main

import (
	"app/domain"
	"app/persistence"
	"app/service"
	"fmt"
)

const xmlFilePath = "users.xml"

func main() {
	userRepo := persistence.NewFileUserRepository(xmlFilePath)
	userService := service.NewUserService(userRepo)

	// Simulate adding a user
	err := userService.AddUser(domain.User{ID: "1", Name: "John Doe", Email: "john.doe@example.com"})
	if err != nil {
		fmt.Println("Error saving user:", err)
	}

	// List all users
	users, err := userService.ListUsers()
	if err != nil {
		fmt.Println("Error loading users:", err)
	}

	for _, user := range users {
		fmt.Printf("User: %+v\n", user)
	}
}
