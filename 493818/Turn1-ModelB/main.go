package main

func SumEvenInts(numbers []int) int {
	total := 0
	for _, num := range numbers {
		if num%2 == 0 {
			total += num
		}
	}
	return total
}
