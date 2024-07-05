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
	expect(piscine.IterativeFactorial(0), 1, t)
	expect(piscine.IterativeFactorial(1), 1, t)
	expect(piscine.IterativeFactorial(2), 2, t)
	expect(piscine.IterativeFactorial(3), 6, t)
	expect(piscine.IterativeFactorial(4), 24, t)
	expect(piscine.IterativeFactorial(-5), 0, t)
	expect(piscine.IterativeFactorial(15), 1307674368000, t)
	expect(piscine.IterativeFactorial(16), 20922789888000, t)
	expect(piscine.IterativeFactorial(17), 355687428096000, t)
	expect(piscine.IterativeFactorial(19), 121645100408832000, t)
	expect(piscine.IterativeFactorial(20), 2432902008176640000, t)
	expect(piscine.IterativeFactorial(-20), 0, t)
	expect(piscine.IterativeFactorial(21), 0, t)
}
