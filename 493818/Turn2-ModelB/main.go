package main

// SumEvenIntegers returns the sum of all even integers in the input array.
func SumEvenIntegers(arr []int) int {
	sum := 0
	for _, num := range arr {
		if num%2 == 0 {
			sum += num
		}
	}
	return sum
}
