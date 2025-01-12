package main

import (
	"testing"
)

func TestSumEvenIntegersParallel(t *testing.T) {
	// Test cases for the parallelized function
	testCases := []struct {
		name   string
		arr    []int
		expect int
	}{
		{"empty array", []int{}, 0},
		{"single even number", []int{2}, 2},
		{"single odd number", []int{1}, 0},
		{"mixed numbers", []int{1, 2, 3, 4, 5}, 6},
		{"large array with even sum", []int{1, 2, 3, 4}, 6},
		{"large array with odd sum", []int{1, 3, 5, 7}, 0},
		{"large array with all even", []int{2, 4, 6, 8, 10}, 30},
		{"large array with all odd", []int{1, 3, 5, 7, 9}, 0},
		{"very large array", make([]int, 1000000, 1000000), 0}, // Fill with even/odd numbers as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SumEvenIntegersParallel(tc.arr)
			if result != tc.expect {
				t.Errorf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func BenchmarkSumEvenIntegersParallel(b *testing.B) {
	// Benchmark the parallelized function
	largeArray := make([]int, 10000)
	for i := 0; i < len(largeArray); i++ {
		largeArray[i] = i % 2 * 10 // Make 50% even numbers
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumEvenIntegersParallel(largeArray)
	}
}
