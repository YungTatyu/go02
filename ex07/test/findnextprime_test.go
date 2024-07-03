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

func TestFindNextPrime(t *testing.T) {
	expect(piscine.FindNextPrime(0), 2, t)
	expect(piscine.FindNextPrime(1), 2, t)
	expect(piscine.FindNextPrime(-3), 2, t)
	expect(piscine.FindNextPrime(4), 5, t)
	expect(piscine.FindNextPrime(9223372036854775807), 9223372036854775807, t)
}
