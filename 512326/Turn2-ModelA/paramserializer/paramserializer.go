package paramserializer

import (
	"fmt"
	"net/url"
	"strconv"
)

// Address represents the nested address structure for a user
type Address struct {
	City  string
	State string
}

// User represents the ORM model for a user with an embedded Address
type User struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	Age     int
	Address Address
}

// SerializeQueryParams serializes URL query parameters into the User struct
func SerializeQueryParams(rawQuery string) (*User, error) {
	params, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil, err
	}

	user := User{}
	if idStr, exists := params["user_id"]; exists {
		if id, err := strconv.Atoi(idStr[0]); err == nil {
			user.ID = id
		} else {
			return nil, fmt.Errorf("invalid user_id: %v", err)
		}
	}

	if name, exists := params["name"]; exists {
		user.Name = name[0]
	}

	if ageStr, exists := params["age"]; exists {
		if age, err := strconv.Atoi(ageStr[0]); err == nil {
			user.Age = age
		} else {
			return nil, fmt.Errorf("invalid age: %v", err)
		}
	}

	// Extract nested address parameters
	if city, exists := params["address[city]"]; exists {
		user.Address.City = city[0]
	}
	if state, exists := params["address[state]"]; exists {
		user.Address.State = state[0]
	}

	return &user, nil
}
