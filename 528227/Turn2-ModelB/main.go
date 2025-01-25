package main

import (
	"fmt"
	"modelb/user"
)

func main() {
	// Create a UserService using the factory function
	userService := user.NewUserService()

	user, err := userService.GetUser(1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("User: %v\n", user)
}
