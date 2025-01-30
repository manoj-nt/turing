package main

import (
	"fmt"
	"os"
)

// Function that simulates resource allocation and cleanup using defer
func OpenAndProcessFile(filePath string) error {
	// Simulate resource allocation (file opening)
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	// Ensure the file is closed when the function exits
	defer file.Close()

	// Simulate some processing on the file
	// If an error occurs during processing, cleanup still happens
	_, err = file.Write([]byte("some data"))
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
