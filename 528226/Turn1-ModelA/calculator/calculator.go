// calculator.go
package calculator

import "fmt"

// Add sums two integers
func Add(a, b int) int {
	return a + b
}

// Subtract subtracts second integer from the first
func Subtract(a, b int) int {
	return a - b
}

// Multiply multiplies two integers
func Multiply(a, b int) int {
	return a * b
}

// Divide divides the first integer by the second
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
