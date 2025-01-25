package user_test

import (
	"modela/user" // Assume this imports the package above
	"testing"
)

// Benchmark using concrete type.
func BenchmarkConcreteType(b *testing.B) {
	service := &user.RealUserService{}
	for i := 0; i < b.N; i++ {
		_, _ = service.GetUser(1)
	}
}

// Benchmark using interface type.
func BenchmarkInterfaceType(b *testing.B) {
	var service user.UserService = &user.RealUserService{}
	for i := 0; i < b.N; i++ {
		_, _ = service.GetUser(1)
	}
}
