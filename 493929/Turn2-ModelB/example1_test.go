package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockProcessItem struct {
	called bool
	err    error
}

func (m *MockProcessItem) ProcessItem(item *Item) (bool, error) {
	m.called = true
	return true, m.err
}

func TestProcessCollection_ManualMocking(t *testing.T) {
	mock := &MockProcessItem{}
	err := ProcessCollection([]Item{{ID: 1}}, mock.ProcessItem)
	assert.NoError(t, err)
	assert.True(t, mock.called)
}
