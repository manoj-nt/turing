package main

// Define the callback function type
type Callback func(int) int

// Function that uses the callback multiple times
func ProcessDataWithCallback(callback Callback, data []int) []int {
	var results []int
	for _, d := range data {
		results = append(results, callback(d))
	}
	return results
}
