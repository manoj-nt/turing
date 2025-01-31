// counter_test.go
package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	c := &Counter{}
	var wg sync.WaitGroup
	goroutines := 100

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	wg.Wait()

	if c.Value() != goroutines {
		t.Errorf("expected %d, got %d", goroutines, c.Value())
	}
}
