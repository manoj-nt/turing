package main

import (
	"testing"
)

func TestProcessCollectionMockery(t *testing.T) {
	items := []Item{{ID: 1, Name: "Item 1"}}

	testCases := []struct {
		name      string
		callback  ProcessItem
		expectErr bool
	}{
		{
			name:      "Process Collection with Mockery",
			callback:  new(mocks.ProcessItem),
			expectErr: false, // All items processed
		},
		{
			name:      "Process Collection with Mockery that Fails",
			callback:  new(mocks.ProcessItem),
			expectErr: true, // Stops on Item 1
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := tc.callback.(*mocks.ProcessItem)
			mock.On("Process", &items[0]).Return(false, nil)
			err := ProcessCollection(items, tc.callback)
			if (err != nil) != tc.expectErr {
				t.Fatalf("Expected error %v, got %v", tc.expectErr, err != nil)
			}
			mock.AssertExpectations(t)
		})
	}
}
