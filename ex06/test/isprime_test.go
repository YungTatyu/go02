package test

import (
	"piscine"
	"testing"
)

func expect(actual, expect bool, t *testing.T) {
	if actual != expect {
		t.Errorf("expected %t, but got %t", expect, actual)
	}
}

func TestIterativeFactorial(t *testing.T) {
	expect(piscine.IsPrime(0), false, t)
	expect(piscine.IsPrime(1), false, t)
	expect(piscine.IsPrime(-3), false, t)
	expect(piscine.IsPrime(4), false, t)
	expect(piscine.IsPrime(2), true, t)
	expect(piscine.IsPrime(3), true, t)
	expect(piscine.IsPrime(5), true, t)
	expect(piscine.IsPrime(7), true, t)
	expect(piscine.IsPrime(11), true, t)
	expect(piscine.IsPrime(13), true, t)
	expect(piscine.IsPrime(17), true, t)
	expect(piscine.IsPrime(19), true, t)
	expect(piscine.IsPrime(23), true, t)
	expect(piscine.IsPrime(29), true, t)
	expect(piscine.IsPrime(31), true, t)
	expect(piscine.IsPrime(37), true, t)
	expect(piscine.IsPrime(41), true, t)
	expect(piscine.IsPrime(43), true, t)
	expect(piscine.IsPrime(9223372036854775807), true, t)
}
