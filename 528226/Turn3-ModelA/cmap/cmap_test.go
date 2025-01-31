// cmap_test.go
package cmap

import (
	"sync"
	"testing"
)

func TestConcurrentMap_SetAndGet(t *testing.T) {
	cmap := NewConcurrentMap()
	var wg sync.WaitGroup

	// Test setting values concurrently.
	wg.Add(1)
	go func() {
		defer wg.Done()
		cmap.Set("key1", 10)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmap.Set("key2", 20)
	}()

	wg.Wait()

	// Test getting values.
	val, err := cmap.Get("key1")
	if err != nil || val != 10 {
		t.Errorf("expected 10, got %d, error: %v", val, err)
	}
	val, err = cmap.Get("key2")
	if err != nil || val != 20 {
		t.Errorf("expected 20, got %d, error: %v", val, err)
	}
}

func TestConcurrentMap_GetError(t *testing.T) {
	cmap := NewConcurrentMap()

	// Test getting a non-existent key.
	_, err := cmap.Get("nonexistent")
	if err == nil || err != ErrKeyNotFound {
		t.Errorf("expected error %v, got %v", ErrKeyNotFound, err)
	}
}
