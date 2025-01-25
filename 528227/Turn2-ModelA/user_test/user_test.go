package user_test

import (
	"testing"

	"modela/user" // Assume this imports user package
)

// TestUserProcessorWithRealService verifies that UserProcessor works with RealUserService.
func TestUserProcessorWithRealService(t *testing.T) {
	service := &user.RealUserService{}
	processor := user.NewUserProcessor(service)

	title, err := processor.ProcessUserTile(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expectedTitle := "Mr. John Doe"
	if title != expectedTitle {
		t.Errorf("expected %s, got %s", expectedTitle, title)
	}
}

// TestUserProcessorWithMockService verifies that UserProcessor works correctly with MockUserService.
func TestUserProcessorWithMockService(t *testing.T) {
	mockData := map[int]*user.User{
		2: {ID: 2, Name: "Jane Doe"},
	}
	mockService := &user.MockUserService{MockData: mockData}
	processor := user.NewUserProcessor(mockService)

	title, err := processor.ProcessUserTile(2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expectedTitle := "Ms. Jane Doe"
	if title != expectedTitle {
		t.Errorf("expected %s, got %s", expectedTitle, title)
	}
}
