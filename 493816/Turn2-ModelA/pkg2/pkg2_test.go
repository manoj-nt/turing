package pkg2

import (
	"testing"
	"time"

	"modela/timeout" // import the timeout package
)

func TestDatabaseConnection(t *testing.T) {
	// Use the timeout management interface
	timeout.ApplyTimeoutToTest("TestDatabaseConnection", func() {
		// Your test logic here
		time.Sleep(3 * time.Second) // Example: Normal test
	})
}

func TestAPIRequest(t *testing.T) {
	// Use the timeout management interface
	timeout.ApplyTimeoutToTest("TestAPIRequest", func() {
		// Your test logic here
		time.Sleep(2 * time.Second) // Example: API test
	})
}
