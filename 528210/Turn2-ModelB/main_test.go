package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestOpenAndProcessFile_DeferClose(t *testing.T) {
	mockFile := new(MockFile)
	// Stub the Close method to keep track of whether it was called or not
	calledClose := false
	mockFile.On("Close").Run(func(args mock.Arguments) {
		calledClose = true
	}).Return(nil)

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

	// Assert that Close was called using the stubbed behavior
	assert.True(t, calledClose, "Close method was not called")
}
