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
	expect(piscine.RecursiveFactorial(0), 1, t)
	expect(piscine.RecursiveFactorial(1), 1, t)
	expect(piscine.RecursiveFactorial(2), 2, t)
	expect(piscine.RecursiveFactorial(3), 6, t)
	expect(piscine.RecursiveFactorial(4), 24, t)
	expect(piscine.RecursiveFactorial(-5), -120, t)
	expect(piscine.RecursiveFactorial(15), 1307674368000, t)
	expect(piscine.RecursiveFactorial(16), 20922789888000, t)
	expect(piscine.RecursiveFactorial(17), 355687428096000, t)
	expect(piscine.RecursiveFactorial(19), 121645100408832000, t)
	expect(piscine.RecursiveFactorial(20), 2432902008176640000, t)
	expect(piscine.RecursiveFactorial(-20), -2432902008176640000, t)
	expect(piscine.RecursiveFactorial(21), 0, t)
}
