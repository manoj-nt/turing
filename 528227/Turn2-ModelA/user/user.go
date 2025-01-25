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

// GetUser simulates retrieving a user by ID.
func (s *RealUserService) GetUser(userID int) (*User, error) {
	return &User{ID: userID, Name: "John Doe"}, nil
}

// UserProcessor processes user data using a UserService.
type UserProcessor struct {
	userService UserService
}

// NewUserProcessor constructs a UserProcessor with the given UserService.
func NewUserProcessor(service UserService) *UserProcessor {
	return &UserProcessor{
		userService: service,
	}
}

// ProcessUserTile generates a user title using its ID.
func (p *UserProcessor) ProcessUserTile(userID int) (string, error) {
	user, err := p.userService.GetUser(userID)
	if err != nil {
		return "", err
	}
	return "Mr. " + user.Name, nil
}

// MockUserService mocks the UserService for testing purposes.
type MockUserService struct {
	MockData map[int]*User
}

// GetUser returns a user if present in the mock data.
func (m *MockUserService) GetUser(userID int) (*User, error) {
	user, exists := m.MockData[userID]
	if !exists {
		return nil, nil
	}
	return user, nil
}
