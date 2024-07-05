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
	expect(piscine.Sqrt(0), 0, t)
	expect(piscine.Sqrt(1), 1, t)
	expect(piscine.Sqrt(-3), 0, t)
	expect(piscine.Sqrt(2), 0, t)
	expect(piscine.Sqrt(4), 2, t)
	expect(piscine.Sqrt(9223372030926249001), 3037000499, t)
	expect(piscine.Sqrt(9223372030926249002), 0, t)
}
