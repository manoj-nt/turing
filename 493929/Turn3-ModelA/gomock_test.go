package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestProcessDataWithGomockMultipleCalls(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock of the Callback interface
	mockCallback := NewMockCallback(ctrl)

	// Set the expectations for the mock callback
	mockCallback.EXPECT().ProcessData("call").Return("Processed call ")
	mockCallback.EXPECT().ProcessData("another call").Return("Processed another call ")

	// Pass the mock callback to the function
	result := ProcessData(mockCallback)

	expected := "Processed call Processed another call "
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
