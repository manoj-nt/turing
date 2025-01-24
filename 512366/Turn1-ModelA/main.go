package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

// Config represents a configuration with its key-value pairs
type Config struct {
	Data map[string]string
	mu   sync.RWMutex // Mutex to handle concurrency safety
}

// LoadConfig loads the configuration from multiple files concurrently
func LoadConfig(files []string) (*Config, error) {
	config := &Config{
		Data: make(map[string]string),
	}

	var wg sync.WaitGroup
	errs := make(chan error, len(files))

	for _, file := range files {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			if err := config.loadFromFile(filename); err != nil {
				errs <- err
			}
		}(file)
	}

	wg.Wait()
	close(errs)

	// Collect any errors
	for err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}

// loadFromFile reads and parses a configuration file
func (c *Config) loadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' { // Skip empty lines and comments
			continue
		}
		parts := splitLine(line)
		if len(parts) == 2 {
			c.mu.Lock()
			c.Data[parts[0]] = parts[1]
			c.mu.Unlock()
		}
	}

	return scanner.Err()
}

// splitLine parses a "key=value" line into two strings
func splitLine(line string) []string {
	for i := 0; i < len(line); i++ {
		if line[i] == '=' {
			return []string{line[:i], line[i+1:]}
		}
	}
	return nil
}

// Get retrieves a configuration value for a given key
func (c *Config) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.Data[key]
	return value, exists
}

func main() {
	// Assume we have config files "config1.txt", "config2.txt"
	files := []string{"config1.txt", "config2.txt"}

	start := time.Now()
	config, err := LoadConfig(files)
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("Configuration loaded in %v\n", elapsed)

	// Example usage
	value, exists := config.Get("someKey")
	if exists {
		fmt.Printf("Value for 'someKey': %s\n", value)
	} else {
		fmt.Println("'someKey' does not exist")
	}
}
