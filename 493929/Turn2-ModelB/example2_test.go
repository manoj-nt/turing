package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProcessItem struct {
	mock.Mock
}

func (m *MockProcessItem) ProcessItem(item *Item) (bool, error) {
	args := m.Called(item)
	return args.Bool(0), args.Error(1)
}

func TestProcessCollection_Mockery(t *testing.T) {
	mock := &MockProcessItem{}
	mock.On("ProcessItem", mock.MatchedBy(func(item *Item) bool {
		return item.ID == 1
	})).Return(true, nil)

	err := ProcessCollection([]Item{{ID: 1}}, mock.ProcessItem)
	assert.NoError(t, err)
	mock.AssertCalled(t, "ProcessItem", &Item{ID: 1})
}
