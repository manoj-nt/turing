package paramserializer

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Coordinates represents the geolocation coordinates
type Coordinates struct {
	Lat float64
	Lng float64
}

// Address represents a user's address with coordinates
type Address struct {
	City        string
	State       string
	Coordinates Coordinates
}

// User represents the ORM model for a user, with an embedded Address struct
type User struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	Age     int
	Tags    []string `gorm:"-"`
	Address Address  `gorm:"embedded"` // GORM handles nested structs
}

// SerializeQueryParams parses a query string and maps the parameters to a struct
func SerializeQueryParams(rawQuery string) (*User, error) {
	params, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil, err
	}

	user := User{}

	// Handle data mapping
	for key, values := range params {
		// Handle nested fields
		if parts := strings.Split(key, "["); len(parts) > 1 {
			field := parts[0]
			subfield := strings.TrimRight(parts[1], "]")

			if len(parts) > 2 { // For deeper nesting
				subSubfield := strings.TrimRight(parts[2], "]")
				// Nested parameter handling for Address' Coordinates
				if field == "address" && subfield == "coordinates" {
					switch subSubfield {
					case "lat":
						if lat, err := strconv.ParseFloat(values[0], 64); err == nil {
							user.Address.Coordinates.Lat = lat
						} else {
							return nil, fmt.Errorf("invalid latitude: %v", err)
						}
					case "lng":
						if lng, err := strconv.ParseFloat(values[0], 64); err == nil {
							user.Address.Coordinates.Lng = lng
						} else {
							return nil, fmt.Errorf("invalid longitude: %v", err)
						}
					}
				}
			} else if field == "address" { // Nested Address fields
				switch subfield {
				case "city":
					user.Address.City = values[0]
				case "state":
					user.Address.State = values[0]
				}
			}
		} else if strings.HasSuffix(key, "[]") { // Handle slice fields like 'tags[]'
			baseKey := strings.TrimSuffix(key, "[]")
			switch baseKey {
			case "tags":
				user.Tags = append(user.Tags, values...)
			}
		} else { // Handle regular fields
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

	return &user, nil
}

// ValidateUser performs validation including for new slice fields and deeper nests
func ValidateUser(user *User) error {
	// Basic validation for User
	if user.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if user.Age <= 0 {
		return fmt.Errorf("age must be greater than zero")
	}
	// Ensure Tags have non-empty values, sanitize as needed
	for i, tag := range user.Tags {
		if strings.TrimSpace(tag) == "" {
			return fmt.Errorf("tag at index %d is empty", i)
		}
		user.Tags[i] = strings.TrimSpace(tag)
	}
	// Validate Address, including Coordinates
	if err := sanitizeAddress(&user.Address); err != nil {
		return fmt.Errorf("address validation failed: %v", err)
	}

	return nil
}

// sanitize provides deeper validation on structures like coordinates
func sanitizeAddress(address *Address) error {
	if address.City == "" {
		return fmt.Errorf("city cannot be empty")
	}

	if address.State == "" {
		return fmt.Errorf("state cannot be empty")
	}

	address.City = strings.TrimSpace(address.City)
	address.State = strings.TrimSpace(address.State)
	// Additional checks (perhaps) for valid Lat/Lng values
	return nil
}
