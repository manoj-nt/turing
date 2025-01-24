package user

// User represents basic user information
type User struct {
	ID   int
	Name string
}

// UserService defines a contract for user operations
type UserService interface {
	GetUser(userID int) (*User, error)
}

// RealUserService is a concrete implementation of UserService
type RealUserService struct{}

// GetUser fetches a user by ID
func (s *RealUserService) GetUser(userID int) (*User, error) {
	// In a real application, this might query a database
	return &User{ID: userID, Name: "John Doe"}, nil
}
