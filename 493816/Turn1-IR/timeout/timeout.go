package timeout

import (
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

// TimeoutForTest applies the timeout for a given test function
func TimeoutForTest(testName string) time.Duration {
	timeout := GetTimeout(testName)
	return timeout
}

// ApplyTimeoutToTest will ensure the test doesn't exceed its timeout duration
func ApplyTimeoutToTest(testName string, testFunc func()) {
	timeout := TimeoutForTest(testName)
	timer := time.NewTimer(timeout)
	done := make(chan bool)

	go func() {
		testFunc()
		done <- true
	}()

	select {
	case <-done:
		// Test finished
		timer.Stop()
	case <-timer.C:
		// Timeout reached
		fmt.Printf("Test %s timed out after %s\n", testName, timeout)
	}
}
