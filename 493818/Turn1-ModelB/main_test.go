package main

import "testing"

func TestSumEvenInts(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Single even element",
			input:    []int{2},
			expected: 2,
		},
		{
			name:     "Single odd element",
			input:    []int{1},
			expected: 0,
		},
		{
			name:     "Even and odd elements",
			input:    []int{2, 3, 4, 5, 6},
			expected: 12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := SumEvenInts(tc.input)
			if actual != tc.expected {
				t.Errorf("Got %d, wanted %d", actual, tc.expected)
			}
		})
	}
}

func BenchmarkSumEvenInts(b *testing.B) {
	testData := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	for i := 0; i < b.N; i++ {
		SumEvenInts(testData)
	}
}
