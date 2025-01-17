package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestProcessCollectionGomock(t *testing.T) {
	items := []Item{{ID: 1, Name: "Item 1"}}

	testCases := []struct {
		name      string
		callback  ProcessItem
		expectErr bool
	}{
		{
			name:      "Process Collection with Gomock",
			callback:  new(mocks.ProcessItem),
			expectErr: false, // All items processed
		},
		{
			name:      "Process Collection with Gomock that Fails",
			callback:  new(mocks.ProcessItem),
			expectErr: true, // Stops on Item 1
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := mocks.NewMockProcessItem(ctrl)
			mock.EXPECT().Process(&items[0]).Return(false, nil)
			mock.EXPECT().Process(&items[0]).Return(false, nil)
			err := ProcessCollection(items, mock)
			if (err != nil) != tc.expectErr {
				t.Fatalf("Expected error %v, got %v", tc.expectErr, err != nil)
			}
		})
	}
}
