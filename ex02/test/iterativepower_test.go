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
	expect(piscine.IterativePower(0, 1), 0, t)
	expect(piscine.IterativePower(1, 5), 1, t)
	expect(piscine.IterativePower(4, 3), 64, t)
	expect(piscine.IterativePower(-4, 3), -64, t)
	expect(piscine.IterativePower(4, -3), 0, t)
	expect(piscine.IterativePower(5, 0), 1, t)
	expect(piscine.IterativePower(0, 0), 1, t)
	expect(piscine.IterativePower(-0, 0), 1, t)
}
