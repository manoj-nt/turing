// calculator_test.go
package calculator

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("expected 3, got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	if result != 2 {
		t.Errorf("expected 2, got %d", result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	if result != 6 {
		t.Errorf("expected 6, got %d", result)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(6, 3)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if result != 2 {
		t.Errorf("expected 2, got %d", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(6, 0)
	if err == nil {
		t.Errorf("expected error, got none")
	}
}
