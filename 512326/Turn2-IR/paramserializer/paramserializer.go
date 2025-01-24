package paramserializer

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Address represents a user's address
type Address struct {
	City  string
	State string
}

// User represents the ORM model for a user, with an embedded Address struct
type User struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	Age     int
	Address Address `gorm:"embedded"` // GORM will embed Address as part of User's table
}

// SerializeQueryParams parses a query string and maps the parameters to a struct
func SerializeQueryParams(rawQuery string) (*User, error) {
	params, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil, err
	}

	user := User{}

	// Iterate over query parameters and map them to the User struct and its nested structs
	for key, values := range params {
		// Handle nested fields (e.g., address[city])
		if strings.Contains(key, "[") && strings.Contains(key, "]") {
			parts := strings.Split(key, "[")
			field := parts[0]
			subfield := strings.TrimRight(parts[1], "]")

			// Map nested fields to their respective structs
			if field == "address" {
				switch subfield {
				case "city":
					user.Address.City = values[0]
				case "state":
					user.Address.State = values[0]
				}
			}
		} else {
			// Handle regular fields (e.g., user_id, name)
			switch key {
			case "user_id":
				if id, err := strconv.Atoi(values[0]); err == nil {
					user.ID = id
				} else {
					return nil, fmt.Errorf("invalid user_id: %v", err)
				}
			case "name":
				user.Name = values[0]
			case "age":
				if age, err := strconv.Atoi(values[0]); err == nil {
					user.Age = age
				} else {
					return nil, fmt.Errorf("invalid age: %v", err)
				}
			}
		}
	}

	// Return the populated user object
	return &user, nil
}

// sanitizeAddress ensures that the address is valid and sanitized
func sanitizeAddress(address *Address) error {
	// Check that city and state are not empty
	if address.City == "" {
		return fmt.Errorf("city cannot be empty")
	}
	if address.State == "" {
		return fmt.Errorf("state cannot be empty")
	}
	// Further sanitization could be done here (e.g., trim spaces, lowercase, etc.)
	address.City = strings.TrimSpace(address.City)
	address.State = strings.TrimSpace(address.State)

	return nil
}

// ValidateUser performs additional validation on the User struct
func ValidateUser(user *User) error {
	// Basic validation for the user fields
	if user.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if user.Age <= 0 {
		return fmt.Errorf("age must be greater than zero")
	}

	// Validate address fields
	if err := sanitizeAddress(&user.Address); err != nil {
		return fmt.Errorf("address validation failed: %v", err)
	}

	return nil
}
