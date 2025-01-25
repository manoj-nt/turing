package user

// User represents basic user information.
type User struct {
	ID   int
	Name string
}

// UserService defines a contract for user operations.
type UserService interface {
	GetUser(userID int) (*User, error)
}

// RealUserService is a concrete implementation of UserService.
type RealUserService struct{}

// GetUser fetches a user by ID (simulated database query).
func (s *RealUserService) GetUser(userID int) (*User, error) {
	// Simulate a database query or external API call.
	return &User{ID: userID, Name: "John Doe"}, nil
}

// MockUserService is a mock implementation of UserService for testing.
type MockUserService struct {
	MockData map[int]*User
}

// GetUser returns a mocked user based on predefined data.
func (m *MockUserService) GetUser(userID int) (*User, error) {
	user, exists := m.MockData[userID]
	if !exists {
		return nil, nil
	}
	return user, nil
}

// NewUserService returns a new UserService instance.
// By default, it returns a RealUserService.
func NewUserService() UserService {
	return &MockUserService{}
}
