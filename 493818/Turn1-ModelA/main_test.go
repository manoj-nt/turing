package main

import "testing"

func TestSumEvenIntegers(t *testing.T) {
	// Test empty array
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SumEvenIntegers(tc.arr)
			if result != tc.expect {
				t.Errorf("expected %d, got %d", tc.expect, result)
			}
		})
	}
}

func BenchmarkSumEvenIntegers(b *testing.B) {
	// Large test array
	largeArray := make([]int, 10000)
	for i := 0; i < len(largeArray); i++ {
		largeArray[i] = i % 2 * 10 // Make 50% even numbers
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumEvenIntegers(largeArray)
	}
}
