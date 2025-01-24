package main

import (
	"testing"
)

func TestFileProcessor_ProcessFile(t *testing.T) {
	// Setup mock file handler
	mockHandler := NewMockFileHandler()
	fileProcessor := NewFileProcessor(mockHandler)

	// Test file does not exist and is created
	err := fileProcessor.ProcessFile("example.txt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check if the file was successfully "written" to
	if !mockHandler.Exists("example.txt") {
		t.Error("expected file to exist after processing")
	}

	// Check the contents of the file
	expectedContent := []byte("Hello, World!")
	if string(mockHandler.files["example.txt"]) != string(expectedContent) {
		t.Errorf("expected file content %s, got %s", expectedContent, mockHandler.files["example.txt"])
	}
}
