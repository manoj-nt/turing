package user

// MockUserService is a mock implementation for testing
type MockUserService struct {
	mockData map[int]*User
}

// GetUser returns a mocked user based on predefined data
func (m *MockUserService) GetUser(userID int) (*User, error) {
	user, exists := m.mockData[userID]
	if !exists {
		return nil, nil
	}
	return user, nil
}
