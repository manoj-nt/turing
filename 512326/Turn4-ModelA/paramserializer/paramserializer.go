package paramserializer

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Coordinates represents geolocation data.
type Coordinates struct {
	Lat float64
	Lng float64
}

// Address represents a user's address with nested coordinates.
type Address struct {
	City        string
	State       string
	Coordinates Coordinates
}

// User represents the ORM model for a user, with embedded Address, Tags slice, and Metadata map.
type User struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Age      int
	Tags     []string          `gorm:"-"`
	Address  Address           `gorm:"embedded"`
	Metadata map[string]string `gorm:"-"`
}

// SerializeQueryParams parses a query string and maps the parameters to a User struct, with support for maps and default values.
func SerializeQueryParams(rawQuery string) (*User, error) {
	params, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil, err
	}

	// Initialize user with default values for optional fields
	user := User{
		Name:     "Default Name", // Default value for Name
		Age:      20,             // Default value for Age
		Tags:     []string{},
		Metadata: make(map[string]string),
		Address: Address{
			City:  "Default City",  // Default value for City
			State: "Default State", // Default value for State
		},
	}

	for key, values := range params {
		if strings.HasSuffix(key, "[]") { // Handle slices like `tags[]`
			baseKey := strings.TrimSuffix(key, "[]")
			switch baseKey {
			case "tags":
				user.Tags = append(user.Tags, values...)
			}
		} else if strings.Contains(key, "[") { // Handle nested and map fields
			parts := strings.Split(key, "[")
			field := parts[0]
			subfield := strings.TrimRight(parts[1], "]")

			if field == "metadata" { // Handle map fields like `metadata[key1]`
				if len(parts) == 2 {
					user.Metadata[subfield] = values[0]
				}
			} else if len(parts) > 2 { // Handle deeper nesting like `address[coordinates][lat]`
				subSubfield := strings.TrimRight(parts[2], "]")
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
			} else if field == "address" {
				switch subfield {
				case "city":
					user.Address.City = values[0]
				case "state":
					user.Address.State = values[0]
				}
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

// ValidateUser validates the User struct, ensuring maps and slices are populated correctly, and sets defaults.
func ValidateUser(user *User) error {
	if strings.TrimSpace(user.Name) == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if user.Age <= 0 {
		return fmt.Errorf("age must be greater than zero")
	}

	for i, tag := range user.Tags {
		user.Tags[i] = strings.TrimSpace(tag)
		if user.Tags[i] == "" {
			return fmt.Errorf("tag at index %d is empty", i)
		}
	}

	if strings.TrimSpace(user.Address.City) == "" {
		return fmt.Errorf("address city cannot be empty")
	}
	if strings.TrimSpace(user.Address.State) == "" {
		return fmt.Errorf("address state cannot be empty")
	}

	// Optionally validate Metadata keys/values if there are specific requirements

	return nil
}
