package timeout

import (
	"fmt"
	"sync"
	"time"
)

var (
	defaultTimeout = 5 * time.Second
	mutex          sync.Mutex // For synchronizing access to timers and done channel
)

// SetDefaultTimeout allows setting the global default timeout
func SetDefaultTimeout(d time.Duration) {
	mutex.Lock()
	defaultTimeout = d
	mutex.Unlock()
}

// GetTimeout retrieves the timeout for a specific package or test
func GetTimeout(testName string) time.Duration {
	testTimeouts := map[string]time.Duration{
		"TestLongRunning": 10 * time.Second,
		"TestFast":        1 * time.Second,
	}

	mutex.Lock()
	timeout := defaultTimeout
	if t, ok := testTimeouts[testName]; ok {
		timeout = t
	}
	mutex.Unlock()

	return timeout
}

// TimeoutForTest applies the timeout for a given test function
func TimeoutForTest(testName string) time.Duration {
	timeout := GetTimeout(testName)
	// Add a small buffer to account for timer delays
	return timeout + 100*time.Millisecond
}

// ApplyTimeoutToTest ensures that the test doesn't exceed its timeout duration
func ApplyTimeoutToTest(testName string, testFunc func()) {
	timeout := TimeoutForTest(testName)
	timer := time.NewTimer(timeout)
	done := make(chan bool)

	mutex.Lock()
	mutex.Unlock()

	go func() {
		defer func() {
			mutex.Lock()
			done <- true
			mutex.Unlock()
		}()
		testFunc()
	}()

	select {
	case <-done:
		// Test finished successfully, stop the timer
		if !timer.Stop() {
			<-timer.C // Drain the channel to avoid a goroutine leak
		}
	case <-timer.C:
		// Timeout reached
		mutex.Lock()
		fmt.Printf("Test %s timed out after %s\n", testName, timeout)
		mutex.Unlock()
	}
}
