// Package timeoutmanagement provides a unified interface for managing test timeouts.
package timeoutmanagement

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// TimeoutConfig contains the configuration for a specific test timeout.
type TimeoutConfig struct {
	Name    string
	Timeout time.Duration
}

// TimeoutManager manages the timeouts for various test operations.
type TimeoutManager struct {
	configs map[string]TimeoutConfig
	mu      sync.Mutex
}

// NewTimeoutManager returns a new TimeoutManager instance.
func NewTimeoutManager() *TimeoutManager {
	return &TimeoutManager{configs: make(map[string]TimeoutConfig)}
}

// RegisterTimeout registers a new timeout configuration for a test.
func (tm *TimeoutManager) RegisterTimeout(config TimeoutConfig) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	tm.configs[config.Name] = config
}

// GetTimeout retrieves the timeout configuration for a specified test.
func (tm *TimeoutManager) GetTimeout(name string) (TimeoutConfig, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	config, ok := tm.configs[name]
	if !ok {
		return TimeoutConfig{}, fmt.Errorf("timeout not registered for %q", name)
	}
	return config, nil
}

// ExecuteWithTimeout runs the provided function with the specified timeout.
func (tm *TimeoutManager) ExecuteWithTimeout(name string, f func()) {
	config, err := tm.GetTimeout(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, cancel := time.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	go func() {
		if err := f(); err != nil {
			fmt.Printf("Test %q failed: %v\n", name, err)
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("Test %q timed out after %s\n", name, config.Timeout)
	default:
	}
}
