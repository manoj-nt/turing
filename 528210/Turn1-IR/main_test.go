package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockFile is a mock of the os.File type
type MockFile struct {
	mock.Mock
}

func (m *MockFile) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockFile) Write(b []byte) (n int, err error) {
	args := m.Called(b)
	return args.Int(0), args.Error(1)
}

// OpenAndProcessFileModified takes in a file interface so it can be tested with mocks
func OpenAndProcessFileModified(filePath string, fileOpener func(string) (*MockFile, error)) error {
	// Simulate resource allocation (file opening)
	file, err := fileOpener(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	// Ensure the file is closed when the function exits
	defer file.Close()

	// Simulate some processing on the file
	_, err = file.Write([]byte("some data"))
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func TestOpenAndProcessFile_Success(t *testing.T) {
	mockFile := new(MockFile)
	mockFile.On("Close").Return(nil)
	mockFile.On("Write", []byte("some data")).Return(len("some data"), nil)

	// Use a mock file opener
	fileOpener := func(filePath string) (*MockFile, error) {
		return mockFile, nil
	}

	// Run the function
	err := OpenAndProcessFileModified("test.txt", fileOpener)

	// Check that there was no error and that Close was called
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	mockFile.AssertExpectations(t)
}

func TestOpenAndProcessFile_OpenError(t *testing.T) {
	// Simulate an error when opening the file
	fileOpener := func(filePath string) (*MockFile, error) {
		return nil, errors.New("unable to open file")
	}

	// Run the function
	err := OpenAndProcessFileModified("test.txt", fileOpener)

	// Check that the error is handled correctly
	if err == nil {
		t.Fatal("expected an error, got none")
	}
	if err.Error() != "failed to open file: unable to open file" {
		t.Fatalf("expected 'failed to open file: unable to open file', got %v", err)
	}
}

func TestOpenAndProcessFile_WriteError(t *testing.T) {
	mockFile := new(MockFile)
	mockFile.On("Close").Return(nil)
	mockFile.On("Write", []byte("some data")).Return(0, errors.New("write error"))

	// Simulate successful file opening
	fileOpener := func(filePath string) (*MockFile, error) {
		return mockFile, nil
	}

	// Run the function
	err := OpenAndProcessFileModified("test.txt", fileOpener)

	// Check that the write error is correctly handled
	if err == nil {
		t.Fatal("expected an error, got none")
	}
	if err.Error() != "failed to write to file: write error" {
		t.Fatalf("expected 'failed to write to file: write error', got %v", err)
	}

	// Ensure that Close was still called
	mockFile.AssertExpectations(t)
}

func TestOpenAndProcessFile_DeferClose(t *testing.T) {
	mockFile := new(MockFile)
	mockFile.On("Close").Return(nil)

	// Simulate successful file opening and writing
	mockFile.On("Write", []byte("some data")).Return(len("some data"), nil)

	// Simulate an error during file processing
	fileOpener := func(filePath string) (*MockFile, error) {
		return mockFile, nil
	}

	// Run the function with a simulated error (after Write)
	err := OpenAndProcessFileModified("test.txt", fileOpener)

	// Verify that defer cleanup (Close) was still called even after error
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	mockFile.AssertExpectations(t)
}
