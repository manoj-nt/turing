package user_test

import (
	"testing"

	"modela/user" // Assume this path imports the package above
)

// TestGetUserWithRealService tests RealUserService using actual logic
func TestGetUserWithRealService(t *testing.T) {
	service := &user.RealUserService{}
	u, err := service.GetUser(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if u == nil || u.ID != 1 {
		t.Errorf("expected user with ID 1, got %v", u)
	}
}

// TestGetUserWithMockService tests using the MockUserService
func TestGetUserWithMockService(t *testing.T) {
	mockData := map[int]*user.User{
		2: {ID: 2, Name: "Jane Doe"},
	}
	mockService := &user.MockUserService{MockData: mockData}
	u, err := mockService.GetUser(2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if u == nil || u.ID != 2 {
		t.Errorf("expected user with ID 2, got %v", u)
	}
}
