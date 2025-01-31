// counter.go
package counter

import (
	"errors"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count map[string]int
}

func NewCounter() *Counter {
	return &Counter{
		count: make(map[string]int),
	}
}

func (c *Counter) Increment(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if key == "" {
		return errors.New("empty key is not allowed")
	}

	c.count[key]++
	return nil
}

func (c *Counter) Decrement(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if key == "" {
		return errors.New("empty key is not allowed")
	}

	if val, ok := c.count[key]; ok && val > 0 {
		c.count[key]--
	} else {
		return errors.New("key not found or count is zero")
	}
	return nil
}

func (c *Counter) GetCount(key string) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if key == "" {
		return 0, errors.New("empty key is not allowed")
	}

	val, ok := c.count[key]
	if !ok {
		return 0, errors.New("key not found")
	}
	return val, nil
}
