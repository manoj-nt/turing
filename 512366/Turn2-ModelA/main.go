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
	Data map[string]string
	mu   sync.Mutex // Mutex to handle concurrency safety when accessing the map
}

// LoadConfig function loads configurations from multiple files, with later files overriding earlier ones
func LoadConfig(files []string) (*Config, error) {
	config := &Config{
		Data: make(map[string]string),
	}

	errs := make(chan error, len(files))                // Buffered channel to collect errors
	results := make(chan map[string]string, len(files)) // Channel to collect map results from each file

	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			fileMap, err := parseConfigFile(filename)
			if err != nil {
				errs <- err
				return
			}
			results <- fileMap
		}(file)
	}

	wg.Wait()      // Wait for all goroutines to finish processing the files
	close(errs)    // Close the error channel
	close(results) // Close the results channel

	// Collect any errors
	var collectedError error
	for err := range errs {
		if collectedError == nil {
			collectedError = err
		}
	}

	// Integrate results in order
	for resultMap := range results {
		config.mu.Lock()
		for key, value := range resultMap {
			config.Data[key] = value
		}
		config.mu.Unlock()
	}

	if collectedError != nil {
		return nil, collectedError
	}

	return config, nil
}

// parseConfigFile reads and parses the configuration file into a map
func parseConfigFile(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %v", filename, err)
	}
	defer file.Close()

	fileMap := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' { // Skip empty lines or comments
			continue
		}

		parts := splitLine(line)
		if len(parts) == 2 {
			fileMap[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", filename, err)
	}

	return fileMap, nil
}

// splitLine parses a key-value pair (key=value) from a line
func splitLine(line string) []string {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) == 2 {
		return []string{strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])}
	}
	return nil
}

// Get retrieves a configuration value for a given key
func (c *Config) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, exists := c.Data[key]
	return value, exists
}

func main() {
	files := []string{"config1.txt", "config2.txt"}

	start := time.Now()
	config, err := LoadConfig(files)
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("Configuration loaded in %v\n", elapsed)

	key := "app_name"
	if value, exists := config.Get(key); exists {
		fmt.Printf("Value for '%s': %s\n", key, value)
	} else {
		fmt.Printf("Key '%s' not found\n", key)
	}
}
