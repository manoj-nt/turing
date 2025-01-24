package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// Config struct represents the configuration with key-value pairs
type Config struct {
	Data           map[string]string
	mu             sync.RWMutex // Mutex to handle concurrency safety when accessing the map
	filePriorities []string     // List of files in priority order for merging
}

// LoadConfig function loads configurations from multiple files concurrently
func LoadConfig(files []string, priorities []string) (*Config, error) {
	config := &Config{
		Data:           make(map[string]string),
		filePriorities: priorities,
	}

	var wg sync.WaitGroup
	errs := make(chan error, len(files)) // Buffered channel to collect errors

	for _, file := range files {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			if err := config.loadFromFile(filename); err != nil {
				errs <- err
			}
		}(file)
	}

	wg.Wait()   // Wait for all goroutines to finish
	close(errs) // Close the error channel

	// Collect any errors
	var collectedError error
	for err := range errs {
		if collectedError == nil {
			collectedError = err // Capture the first error
		}
	}
	if collectedError != nil {
		return nil, collectedError // Return the first encountered error
	}

	return config, nil
}

// loadFromFile reads and parses the configuration file
func (c *Config) loadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		parts := splitLine(line)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			// Check if the key exists in the configuration and if the current file has a higher priority than the existing file
			if existingValue, ok := c.Get(key); ok {
				if filePriorityIndex(c.filePriorities, filename) < filePriorityIndex(c.filePriorities, existingValue) {
					continue
				}
			}
			c.mu.Lock() // Lock the map for writing
			c.Data[key] = value
			c.mu.Unlock() // Unlock after writing
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file %s: %v", filename, err)
	}

	return nil
}

// Get retrieves a configuration value for a given key
func (c *Config) Get(key string) (string, bool) {
	c.mu.RLock() // Lock for reading
	defer c.mu.RUnlock()
	value, exists := c.Data[key]
	return value, exists
}

// splitLine parses a key-value pair (key=value) from a line
func splitLine(line string) []string {
	// Trim spaces and split by the first '='
	parts := strings.SplitN(line, "=", 2)
	if len(parts) == 2 {
		return []string{strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])}
	}
	return nil
}

// filePriorityIndex returns the index of a file in the priority list
func filePriorityIndex(priorities []string, filename string) int {
	for index, priorityFile := range priorities {
		if priorityFile == filename {
			return index
		}
	}
	return -1 // If file not found, return -1
}

func main() {
	// Sample configuration files
	files := []string{"config1.txt", "config2.txt", "config3.txt"}

	// Define the priority order for files, with later files overriding earlier ones
	filePriorities := []string{"config3.txt", "config1.txt", "config2.txt"}

	// Start the process of loading configuration files concurrently
	start := time.Now()
	config, err := LoadConfig(files, filePriorities)
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("Configuration loaded in %v\n", elapsed)

	// Example usage: Retrieve a configuration value
	key := "app_name"
	if value, exists := config.Get(key); exists {
		fmt.Printf("Value for '%s': %s\n", key, value)
	} else {
		fmt.Printf("Key '%s' not found\n", key)
	}
}
