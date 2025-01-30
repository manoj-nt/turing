package main

import (
	"os"
	"regexp"
	"testing"
)

// Helper function to create a temporary file for testing
func createTempFile(t *testing.T, content string) string {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "test_file_")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	_, err = tmpfile.WriteString(content)
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpfile.Close()
	return tmpfile.Name()
}

// TestOpenAndProcessFileSuccess ensures the function works correctly with a valid file
func TestOpenAndProcessFileSuccess(t *testing.T) {
	// Create a temporary file for testing
	tmpfile := createTempFile(t, "test data")
	defer os.Remove(tmpfile)

	err := OpenAndProcessFile(tmpfile)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	// Check if the file content was updated as expected
	updatedContent, err := os.ReadFile(tmpfile)
	if err != nil {
		t.Errorf("failed to read updated file content: %v", err)
	}
	expectedContent := "test data"
	if string(updatedContent) != expectedContent {
		t.Errorf("expected file content to be %q, got: %q", expectedContent, string(updatedContent))
	}
}

// TestOpenAndProcessFileFailsToOpen ensures the function handles a file that cannot be opened
func TestOpenAndProcessFileFailsToOpen(t *testing.T) {
	nonExistentFilePath := "non_existent_file.txt"

	err := OpenAndProcessFile(nonExistentFilePath)
	if err == nil {
		t.Error("expected error when opening non-existent file, got none")
	}

	expectedError := "open non_existent_file.txt: no such file or directory"
	if err.Error() != expectedError {
		t.Errorf("expected error %q, got: %q", expectedError, err.Error())
	}
}

// TestOpenAndProcessFileFailsToWrite ensures the function handles an error during file writing
func TestOpenAndProcessFileFailsToWrite(t *testing.T) {
	// Create a temporary file for testing, but make it read-only
	tmpfile := createTempFile(t, "test data")
	defer os.Remove(tmpfile)

	err := os.Chmod(tmpfile, 0400) // Make the file read-only
	if err != nil {
		t.Fatalf("failed to set read-only permissions on temp file: %v", err)
	}

	err = OpenAndProcessFile(tmpfile)
	if err == nil {
		t.Error("expected error when writing to read-only file, got none")
	}

	expectedError := "write test_file_[0-9a-fA-F]+: permission denied"
	if !regexp.MustCompile(expectedError).MatchString(err.Error()) {
		t.Errorf("expected error %q, got: %q", expectedError, err.Error())
	}
}
