package main

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"
)

// Assuming CustomReader and CustomWriter are already defined as in the previous example

// TestCustomReader tests the CustomReader.
func TestCustomReader(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		bufferSize   int
		expectedRead string
		expectedErr  error
	}{
		{"Normal case", "Hello, World!", 5, "Hello, World!", nil},
		{"Empty data", "", 5, "", io.EOF},
		{"Partial read", "Go is great!", 4, "Go is great!", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := NewCustomReader(tt.input)
			buffer := make([]byte, tt.bufferSize)
			var output []byte

			for {
				n, err := reader.Read(buffer)
				output = append(output, buffer[:n]...)
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
			}

			if string(output) != tt.expectedRead {
				t.Errorf("Expected %s, got %s", tt.expectedRead, output)
			}
		})
	}
}

// TestCustomWriter tests the CustomWriter.
func TestCustomWriter(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedErr error
	}{
		{"Normal write", "Hello, Gopher!", nil},
		{"Empty write", "", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file for testing
			tempFile, err := os.CreateTemp("", "testwriter")
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer os.Remove(tempFile.Name()) // Clean up

			writer := &CustomWriter{file: tempFile}

			_, err = writer.Write([]byte(tt.input))
			if !errors.Is(err, tt.expectedErr) {
				t.Fatalf("Expected error %v, got %v", tt.expectedErr, err)
			}

			writer.Close()

			// Verify file content matches input
			content, err := os.ReadFile(tempFile.Name())
			if err != nil {
				t.Fatalf("Failed to read temp file: %v", err)
			}
			if !bytes.Equal(content, []byte(tt.input)) {
				t.Errorf("Expected file content %s, got %s", tt.input, content)
			}
		})
	}
}

// Test error handling for CustomWriter
func TestCustomWriterErrorHandling(t *testing.T) {
	// Use invalid file path to simulate write error.
	invalidPath := string([]byte{0})

	writer, err := NewCustomWriter(invalidPath)
	if err == nil {
		t.Fatal("Expected error when creating writer with invalid path, got nil")
	}
	if writer != nil {
		t.Fatal("Writer should be nil when there's an error opening the file")
	}
}
