package main

import (
	"sync"
)

// SumEvenIntegersParallel returns the sum of all even integers in the input array,
// using parallelism to improve performance for large arrays.
func SumEvenIntegersParallel(arr []int) int {
	const chunkSize = 1000 // Adjust this size based on the system and array characteristics
	var wg sync.WaitGroup
	var partialSums []int

	// Split the array into chunks
	for i := 0; i < len(arr); i += chunkSize {
		wg.Add(1)
		go func(start int, end int) {
			defer wg.Done()
			partialSum := 0
			for j := start; j < end; j++ {
				if arr[j]%2 == 0 {
					partialSum += arr[j]
				}
			}
			partialSums = append(partialSums, partialSum)
		}(i, min(i+chunkSize, len(arr)))
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Sum all partial sums
	totalSum := 0
	for _, partialSum := range partialSums {
		totalSum += partialSum
	}

	return totalSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
