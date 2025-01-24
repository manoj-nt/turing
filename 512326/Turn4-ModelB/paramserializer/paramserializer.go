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

// Define a custom error type for optional field parsing errors
type OptionalFieldError struct {
	msg string
}

func (e OptionalFieldError) Error() string {
	return e.msg
}

// User represents the ORM model for a user, including embedded Address, Tags slice, and Metadata map.
type User struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Age          int
	Tags         []string          `gorm:"-"` // Ignored by GORM
	Address      Address           `gorm:"embedded"`
	Metadata     map[string]string // Map with arbitrary keys
	OptionalBool *bool             `gorm:"-"` // Example of an optional field with default value
}

func (u *User) initDefaults() {
	// Initialize default values for optional fields
	u.OptionalBool = new(bool)
	*u.OptionalBool = true // Set a default value for OptionalBool
}

// SerializeQueryParams parses a query string and maps the parameters to a User struct.
func SerializeQueryParams(rawQuery string) (*User, error) {
	params, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil, err
	}

	user := User{}
	user.initDefaults() // Initialize default values for optional fields

	for key, values := range params {
		if strings.HasSuffix(key, "[]") { // Handle slices like `tags[]`
			baseKey := strings.TrimSuffix(key, "[]")
			switch baseKey {
			case "tags":
				user.Tags = append(user.Tags, values...)
			}
		} else if strings.Contains(key, "[") && strings.Contains(key, "]") { // Handle nested maps like `metadata[key]`
			parts := strings.Split(key, "[")
			field := parts[0]
			subfield := strings.TrimRight(parts[1], "]")

			if field == "metadata" {
				if user.Metadata == nil {
					user.Metadata = make(map[string]string)
				}
				user.Metadata[subfield] = values[0]

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
