package main

import (
	"testing"
)

func Test(t *testing.T) {
	// given
	expected := "issip"
	text := "mississippi"
	pattern := "issip"
	m := len(pattern)

	// when
	res := Dfa("mississippi", "issip")
	found := text[res : res+m]

	// then
	if expected != found {
		t.Errorf("Result is: %v, but expected: %v\n", found, expected)
	}
}
