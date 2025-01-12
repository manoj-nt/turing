package testpackage1

import (
	"testing"
	"time"

	"modela/timeoutmanagement"
)

var tm = timeoutmanagement.NewTimeoutManager()

func TestSlowFunction(t *testing.T) {
	tm.RegisterTimeout(timeoutmanagement.TimeoutConfig{Name: "SlowFunction", Timeout: 10 * time.Second})

	tm.ExecuteWithTimeout("SlowFunction", func() {
		// Simulate a slow operation
		time.Sleep(5 * time.Second)
		t.Log("Slow function completed")
	})

	t.Run("Example", func(t *testing.T) {
		t.Log("Running example")
	})
}
