package gcd

import (
	"testing"
)

func TestIter(t *testing.T) {
	// given
	a := 28851538
	b := 1183019
	expect := 17657

	// when
	r := gcdIter(a, b)

	// then
	if r != expect {
		t.Errorf("Wrong GCD answer. Expected: %d, Given: %d", expect, r)
	}
}

func TestRec(t *testing.T) {
	// given
	a := 28851538
	b := 1183019
	expect := 17657

	// when
	r := gcdRec(a, b)

	// then
	if r != expect {
		t.Errorf("Wrong GCD answer. Expected: %d, Given: %d", expect, r)
	}
}
