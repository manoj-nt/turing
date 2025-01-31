// counter_test.go
package counter

import (
	"errors"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	c := NewCounter()

	// Test basic increment
	err := c.Increment("test")
	if err != nil {
		t.Errorf("Unexpected error on increment: %v", err)
	}
	if val, _ := c.GetCount("test"); val != 1 {
		t.Errorf("Expected count to be 1, got: %d", val)
	}

	// Test edge case: Empty key
	err = c.Increment("")
	if !errors.Is(err, errors.New("empty key is not allowed")) {
		t.Errorf("Expected empty key error, got: %v", err)
	}

	// Test decrement
	err = c.Decrement("test")
	if err != nil {
		t.Errorf("Unexpected error on decrement: %v", err)
	}

	// Test decrement with edge case: Empty key
	err = c.Decrement("")
	if !errors.Is(err, errors.New("empty key is not allowed")) {
		t.Errorf("Expected empty key error, got: %v", err)
	}

	// Test decrement with edge case: Key not found
	err = c.Decrement("non-existent")
	if !errors.Is(err, errors.New("key not found or count is zero")) {
		t.Errorf("Expected key not found error, got: %v", err)
	}

	// Test concurrent access
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := c.Increment("test")
			if err != nil {
				t.Errorf("Unexpected error on concurrent increment: %v", err)
			}
		}()
	}
	wg.Wait()
	val, _ := c.GetCount("test")
	if val != 100 {
		t.Errorf("Expected count to be 100 after concurrent increments, got: %d", val)
	}

}
