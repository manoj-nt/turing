package main

import (
	"modela/mocks" // path to `mocks` package created
	"testing"
)

func TestProcessDataWithMockeryMultipleCalls(t *testing.T) {
	mockCallback := new(mocks.Callback)

	// Set expectations for each call
	mockCallback.On("ProcessData", "call").Return("Processed call ")
	mockCallback.On("ProcessData", "another call").Return("Processed another call ")

	// Call the function
	result := ProcessData(mockCallback)

	expected := "Processed call Processed another call "
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}

	// Assert that all expectations were met
	mockCallback.AssertExpectations(t)
}
