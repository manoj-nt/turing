package pkg1

import (
	"context"
	"testing"
	"time"

	"modela/timeout" // replace it with timeout package's path
)

func TestFast(t *testing.T) {
	// Use the timeout management interface
	timeout.ApplyTimeoutToTest("TestFast", func(ctx context.Context) {
		// Your test logic here
		select {
		case <-time.After(500 * time.Millisecond): // Example: Fast test
		case <-ctx.Done():
			t.Fatal("TestFast timed out")
		}
	})
}

func TestLongRunning(t *testing.T) {
	// Use the timeout management interface
	timeout.ApplyTimeoutToTest("TestLongRunning", func(ctx context.Context) {
		// Your test logic here
		select {
		case <-time.After(8 * time.Second): // Example: Long running test
		case <-ctx.Done():
			t.Fatal("TestLongRunning timed out")
		}
	})
}
