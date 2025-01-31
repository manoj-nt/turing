// calculator_test.go
package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b   int
		expect int
	}{
		{1, 2, 3},
		{-1, -2, -3},
		{0, 0, 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Add(%d, %d)", tc.a, tc.b), func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expect {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expect)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		a, b   int
		expect int
	}{
		{1, 2, 2},
		{-1, -2, 2},
		{0, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Multiply(%d, %d)", tc.a, tc.b), func(t *testing.T) {
			result := Multiply(tc.a, tc.b)
			if result != tc.expect {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expect)
			}
		})
	}
}
