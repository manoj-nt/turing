package main

import (
    "testing"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"
)

type MockProcessItem struct {
    ctrl     *gomock.Controller
    mock     *MockProcessItemMockRecorder
    called   bool
    expected error
}

// MockProcessItemMockRecorder is the mock recorder for MockProcessItem
type MockProcessItemMockRecorder struct {
    mock *MockProcessItem
}

// NewMockProcessItem creates a new mock instance
func NewMockProcessItem(ctrl *gomock.Controller) *MockProcessItem {
    mock := &MockProcessItem{ctrl: ctrl}
    mock.mock = &MockProcessItemMockRecorder{mock}
    return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProcessItem) EXPECT() *MockProcessItemMockRecorder {
    return m.mock
}

// ProcessItem mocks base method
func (m *MockProcessItem) ProcessItem(item *Item) (bool, error) {