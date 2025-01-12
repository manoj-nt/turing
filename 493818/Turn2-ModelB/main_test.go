package main

import "testing"

// OptimizedSumEvenIntegers returns the sum of all even integers in the input array.
func OptimizedSumEvenIntegers(arr []int) int {
	sum := 0
	// Loop through the array with an increment of 2
	for i := 0; i < len(arr); i += 2 {
		// Check if the current index is within the array bounds
		if i < len(arr) {
			sum += arr[i]
		}
	}
	return sum
}

func TestOptimizedSumEvenIntegers(t *testing.T) {
	largeEvenArray := make([]int, 1_000_000)
	for i := 0; i < len(largeEvenArray); i++ {
		largeEvenArray[i] = 2 * i
	}

	expectedSum := 0
	for i := 0; i < len(largeEvenArray); i += 2 {
		expectedSum += largeEvenArray[i]
	}

	result := OptimizedSumEvenIntegers(largeEvenArray)
	if result != expectedSum {
		t.Errorf("OptimizedSumEvenIntegers: expected %d, got %d", expectedSum, result)
	}
}

func BenchmarkSumEvenIntegersLargeArray(b *testing.B) {
	// Array with 10 million elements
	largeArray := make([]int, 10_000_000)
	for i := 0; i < len(largeArray); i++ {
		largeArray[i] = i % 2 * 10 // Make 50% even numbers
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumEvenIntegers(largeArray)
	}
}

func BenchmarkOptimizedSumEvenIntegersLargeArray(b *testing.B) {
	// Reuse the large array from the previous benchmark
	largeArray := make([]int, 10_000_000)
	for i := 0; i < len(largeArray); i++ {
		largeArray[i] = i % 2 * 10 // Make 50% even numbers
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OptimizedSumEvenIntegers(largeArray)
	}
}
