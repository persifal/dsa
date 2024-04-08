package main

import (
	"slices"
	"testing"
)

func Test(t *testing.T) {
}

func check(s, templ string, expected []int, t *testing.T) {
	// when
	res := dfa(s, templ)

	// then
	if slices.Compare(res, expected) != 0 {
		t.Errorf("Result is: %v, but expected: %v\n", res, expected)
	}
}
