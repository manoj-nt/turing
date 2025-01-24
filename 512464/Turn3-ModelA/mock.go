package main

import (
	"errors"
	"os"
)

// MockFileHandler is a mock implementation of the FileHandler interface for testing
type MockFileHandler struct {
	files map[string][]byte
}

// NewMockFileHandler initializes a new MockFileHandler
func NewMockFileHandler() *MockFileHandler {
	return &MockFileHandler{files: make(map[string][]byte)}
}

// Open simulates opening a file
func (m *MockFileHandler) Open(name string) (*os.File, error) {
	if _, exists := m.files[name]; exists {
		return &os.File{}, nil // Return a dummy os.File
	}
	return nil, errors.New("file does not exist")
}

// Close simulates closing a file
func (m *MockFileHandler) Close(file *os.File) error {
	return nil // No-operation
}

// Write simulates writing data to a file
func (m *MockFileHandler) Write(file *os.File, data []byte) (int, error) {
	m.files["example.txt"] = data
	return len(data), nil
}

// Exists checks if the file exists
func (m *MockFileHandler) Exists(name string) bool {
	_, exists := m.files[name]
	return exists
}
