package main

import "testing"

// Callback type
type ProcessDataCallback func(data string) string

// Function that accepts a callback
func ProcessData(callback ProcessDataCallback) string {
	data := "data"
	return callback(data)
}

// Mock callback for testing
func mockCallback(data string) string {
	return "Mocked result: " + data
}

func TestProcessData(t *testing.T) {
	result := ProcessData(mockCallback)

	expected := "Mocked result: data"
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
