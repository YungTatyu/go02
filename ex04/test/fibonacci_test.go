package test

import (
	"piscine"
	"testing"
)

func expect(actual, expect int, t *testing.T) {
	if actual != expect {
		t.Errorf("expected %d, but got %d", expect, actual)
	}
}

func TestIterativeFactorial(t *testing.T) {
	expect(piscine.Fibonacci(0), 0, t)
	expect(piscine.Fibonacci(1), 1, t)
	expect(piscine.Fibonacci(2), 1, t)
	expect(piscine.Fibonacci(20), 6765, t)
	expect(piscine.Fibonacci(-4), -1, t)
	expect(piscine.Fibonacci(92), 7540113804746346429, t)
}
