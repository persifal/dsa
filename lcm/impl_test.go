package gcd

import (
	"testing"
)

func Test(t *testing.T) {
	// given
	a := 761457
	b := 614573
	expect := 467970912861

	// when
	r := lcm(a, b)

	// then
	if r != expect {
		t.Errorf("Wrong LCM answer. Expected: %d, Given: %d", expect, r)
	}
}
