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
	file, err := fileOpener(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	_, err = file.Write([]byte("some data"))
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func TestDeferCloseIsCalled(t *testing.T) {
	mockFile := new(MockFile)
	mockFile.On("Close").Return(nil)

	// Simulate successful file opening and a write error
	mockFile.On("Write", []byte("some data")).Return(0, errors.New("write error"))

	// Simulate successful file opening
	fileOpener := func(filePath string) (*MockFile, error) {
		return mockFile, nil
	}

	// Run the function with a simulated error during writing
	err := OpenAndProcessFileModified("test.txt", fileOpener)

	// Verify that the error is returned as expected
	if err == nil || err.Error() != "failed to write to file: write error" {
		t.Fatalf("expected 'failed to write to file: write error', got %v", err)
	}

	// Assert that Close was called even though writing failed
	mockFile.AssertCalled(t, "Close")

	// Additionally, ensure there are no unexpected calls
	mockFile.AssertExpectations(t)
}
