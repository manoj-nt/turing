package paramserializer

import (
	"fmt"
	"net/url"
	"strconv"
)

// User represents the ORM model for a user
type User struct {
	ID   int `gorm:"primaryKey"`
	Name string
	Age  int
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

	return &user, nil
}
