package main

import (
	"modelir/mocks" // path to `mocks` package created
	"testing"
)

func TestProcessDataWithMockery(t *testing.T) {
	mockCallback := new(mocks.Callback)

	// Set expectation
	mockCallback.On("ProcessData", "data").Return("Mocked result: data")

	// Call the function
	result := ProcessData(mockCallback)

	expected := "Mocked result: data"
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}

	// Assert that the expectation was met
	mockCallback.AssertExpectations(t)
}
