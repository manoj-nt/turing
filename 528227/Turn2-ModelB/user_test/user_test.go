package user_test

import (
	"modelb/user"
	"testing"
)

func TestGetUserWithMockServiceViaFactory(t *testing.T) {
	mockData := map[int]*user.User{
		2: {ID: 2, Name: "Jane Doe"},
	}
	// Create a mock service using the factory function
	mockService := user.NewUserService()
	if mock, ok := mockService.(*user.MockUserService); ok {
		mock.MockData = mockData
	} else {
		t.Fatalf("Unable to set mock data, expected mock type")
	}

	u, err := mockService.GetUser(2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if u == nil || u.ID != 2 {
		t.Errorf("expected user with ID 2, got %v", u)
	}
}

func BenchmarkConcreteTypeViaFactory(b *testing.B) {
	// Use the factory function to create a real service
	service := user.NewUserService()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetUser(1)
	}
}
