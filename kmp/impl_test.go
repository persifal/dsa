package main

import "testing"

func Test(t *testing.T) {
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	res := Kmp(text, pattern)
	if res != 10 {
		t.Fatalf("KMP failed to find match: expecteed 10 != actual %d\n", res)
	}
}
