package main

import (
	"slices"
	"testing"
)

func Test(t *testing.T) {
	// given
	s := "acaabc"
	templ := "aab"
	expected := []int{2}

	// then
	check(s, templ, expected, t)
}

func TestMedium(t *testing.T) {
	// given
	s := "AABAACAADAABAABA"
	templ := "AABA"
	expected := []int{0, 9, 12}

	// then
	check(s, templ, expected, t)
}

func check(s, templ string, expected []int, t *testing.T) {
	// when
	res := rabinKarp(s, templ)

	// then
	if slices.Compare(res, expected) != 0 {
		t.Errorf("Result is: %v, but expected: %v\n", res, expected)
	}
}
