package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockProcessItem is a mock implementation of ProcessItem
func MockProcessItem(item *Item) (bool, error) {
	switch item.ID {
	case 1:
		return true, nil // Successful execution
	case 2:
		return false, nil // Stop processing
	case 3:
		return true, fmt.Errorf("simulated error") // Return an error
	case 4:
		return true, nil // Fails after error
	}
	return true, nil
}

func TestProcessCollection(t *testing.T) {
	items := []Item{
		{ID: 1, Name: "Item 1"},
		{ID: 2, Name: "Item 2"},
		{ID: 3, Name: "Item 3"},
		{ID: 4, Name: "Item 4"},
	}

	testCases := []struct {
		name      string
		callback  ProcessItem
		expectErr bool
	}{
		{
			name:      "Process Collection with Mock",
			callback:  MockProcessItem,
			expectErr: true, // Error on Item 3
		},
		{
			name:      "Successful Processing",
			callback:  func(item *Item) (bool, error) { return true, nil },
			expectErr: false, // All items processed
		},
		{
			name:      "Empty Collection",
			callback:  func(item *Item) (bool, error) { return true, nil },
			expectErr: false, // No items to process
		},
		{
			name:      "Single Item Collection",
			callback:  func(item *Item) (bool, error) { return false, nil },
			expectErr: true, // Stops on Item 1
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ProcessCollection(items, tc.callback)
			assert.Equal(t, tc.expectErr, err != nil, "Unexpected error occurred")
		})
	}
}
