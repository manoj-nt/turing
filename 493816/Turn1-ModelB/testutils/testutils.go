package testutils

import (
	"flag"
	"log"
	"testing"
	"time"
)

var defaultTimeout time.Duration

func init() {
	flag.DurationVar(&defaultTimeout, "test.timeout", 30*time.Second, "default timeout for tests")
	flag.Parse()
}

// TestTimeout sets a timeout for the given test.
func TestTimeout(t *testing.T, timeout time.Duration) {
	if timeout <= 0 {
		timeout = defaultTimeout
	}

	done := make(chan bool)
	go func() {
		timer := time.NewTimer(timeout)
		<-timer.C
		log.Printf("Test timeout exceeded: %v", t.Name())
		t.FailNow()
		done <- true
	}()

	t.Cleanup(func() {
		done <- true
	})
}
