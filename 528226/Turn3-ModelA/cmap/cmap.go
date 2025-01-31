// cmap.go
package cmap

import (
	"errors"
	"sync"
)

// ConcurrentMap is a thread-safe map.
type ConcurrentMap struct {
	mu sync.RWMutex
	m  map[string]int
}

var ErrKeyNotFound = errors.New("key not found")

// NewConcurrentMap creates a new ConcurrentMap.
func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[string]int),
	}
}

// Set adds a key-value pair to the map.
func (c *ConcurrentMap) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

// Get retrieves a value for a key or returns an error if not found.
func (c *ConcurrentMap) Get(key string) (int, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.m[key]
	if !exists {
		return 0, ErrKeyNotFound
	}
	return value, nil
}
