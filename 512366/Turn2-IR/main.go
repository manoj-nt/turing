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
	mu   sync.RWMutex // Mutex to handle concurrency safety when accessing the map
}

// LoadConfig function loads configurations from multiple files, with later files overriding earlier ones
func LoadConfig(files []string) (*Config, error) {
	config := &Config{
		Data: make(map[string]string),
	}

	// Error channel to collect errors from goroutines
	errs := make(chan error, len(files))

	// Process files in order, respecting the provided order of the files slice
	for _, file := range files {
		// Read each file and merge its data into the config
		fileMap, err := parseConfigFile(file)
		if err != nil {
			errs <- err
			continue
		}

		// Locking to ensure sequential writing to the config map, one file at a time
		config.mu.Lock()
		for key, value := range fileMap {
			config.Data[key] = value // Later files overwrite earlier ones
		}
		config.mu.Unlock()
	}

	// Close the error channel
	close(errs)

	// Collect any errors encountered during file reading
	var collectedError error
	for err := range errs {
		if collectedError == nil {
			collectedError = err
		}
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
	c.mu.RLock() // Lock for reading
	defer c.mu.RUnlock()
	value, exists := c.Data[key]
	return value, exists
}

func main() {
	// Sample configuration files
	files := []string{"config1.txt", "config3.txt", "config2.txt"}

	// Start the process of loading configuration files
	start := time.Now()
	config, err := LoadConfig(files)
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
