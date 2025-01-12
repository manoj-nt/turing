package pkg1

import (
	"testing"
	"time"

	"ir/timeout" // import the timeout package
)

func TestFast(t *testing.T) {
	// Use the timeout management interface
	timeout.ApplyTimeoutToTest("TestFast", func() {
		// Your test logic here
		time.Sleep(500 * time.Millisecond) // Example: Fast test
	})
}

func TestLongRunning(t *testing.T) {
	// Use the timeout management interface
	timeout.ApplyTimeoutToTest("TestLongRunning", func() {
		// Your test logic here
		time.Sleep(8 * time.Second) // Example: Long running test
	})
}
