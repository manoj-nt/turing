package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestProcessDataWithGomock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock of the Callback interface
	mockCallback := NewMockCallback(ctrl)

	// Set the expectation for the mock callback
	mockCallback.EXPECT().ProcessData("data").Return("Mocked result: data")

	// Pass the mock callback to the function
	result := ProcessData(mockCallback)

	expected := "Mocked result: data"
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
