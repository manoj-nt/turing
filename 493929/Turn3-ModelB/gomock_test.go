package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProcessDataWithCallback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCallback := NewMockCallback(ctrl)

	// Set up expectations for multiple calls
	mockCallback.EXPECT().Call(1).Return(10)
	mockCallback.EXPECT().Call(2).Return(20)
	mockCallback.EXPECT().Call(3).Return(30)

	data := []int{1, 2, 3}
	results := ProcessDataWithCallback(mockCallback.Call, data)

	expectedResults := []int{10, 20, 30}
	if !assert.Equal(t, expectedResults, results) {
		t.Fail()
	}
}
