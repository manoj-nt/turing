package main

import (
	"modelb/mocks" // path to `mocks` package created
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProcessDataWithCallback1(t *testing.T) {
	mockCallback := new(mocks.Callback)

	// Set up expectations with Run() method
	mockCallback.On("Call", 1).Run(func(args mock.Arguments) {
		mockCallback.Return(10)
	})
	mockCallback.On("Call", 2).Run(func(args mock.Arguments) {
		mockCallback.Return(20)
	})
	mockCallback.On("Call", 3).Run(func(args mock.Arguments) {
		mockCallback.Return(30)
	})

	data := []int{1, 2, 3}
	results := ProcessDataWithCallback(mockCallback.Call, data)

	expectedResults := []int{10, 20, 30}
	if !assert.Equal(t, expectedResults, results) {
		t.Fail()
	}

	// Verify that all expectations were met
	mockCallback.AssertExpectations(t)
}
