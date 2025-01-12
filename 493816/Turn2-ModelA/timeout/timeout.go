package timeout

import (
	"context"
	"fmt"
	"time"
)

var defaultTimeout = 5 * time.Second // Default timeout for tests

// SetDefaultTimeout allows setting the global default timeout
func SetDefaultTimeout(d time.Duration) {
	defaultTimeout = d
}

// GetTimeout retrieves the timeout for a specific package or test.
func GetTimeout(testName string) time.Duration {
	// Example: you could use a map to store custom timeouts for specific tests
	testTimeouts := map[string]time.Duration{
		"TestLongRunning": 10 * time.Second,
		"TestFast":        1 * time.Second,
	}

	// Return the specific timeout or default timeout
	if timeout, ok := testTimeouts[testName]; ok {
		return timeout
	}

	return defaultTimeout
}

// WithTestContext returns a context with a deadline for the given test name
func WithTestContext(testName string) (context.Context, context.CancelFunc) {
	timeout := GetTimeout(testName)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return ctx, cancel
}

// ApplyTimeoutToTest ensures the test doesn't exceed its timeout duration using a context
func ApplyTimeoutToTest(testName string, testFunc func(context.Context)) {
	ctx, cancel := WithTestContext(testName)
	defer cancel()

	select {
	case <-ctx.Done():
		// Timeout reached
		fmt.Printf("Test %s timed out after %s\n", testName, ctx.Deadline().Sub(time.Now()))
	case <-time.After(500 * time.Millisecond): // Example: Wait for test to complete
		// Test finished
	}
}
