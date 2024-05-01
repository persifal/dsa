package main

import (
	du "dsa/utils"
)

func Dfa(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	transitions := transitions(needle)
	q := 0
	for i := 0; i < n; i++ {
		c := int(haystack[i])
		q = transitions[q][c]
		if q == m {
			return i - m + 1
		}
	}

	return -1
}

func transitions(pattern string) [][]int {
	m := len(pattern)
	r := make([][]int, m+1)
	for i := range r {
		r[i] = make([]int, 256)
	}

	for q := 0; q <= m; q++ {
		for c := 0; c < 256; c++ {
			k := du.Min(m, q+1)
			for k > 0 && !isSuffix(pattern[:k], pattern[:q]+string(rune(c))) {
				k--
			}

			r[q][c] = k
		}
	}

	return r
}

func isSuffix(suffix, text string) bool {
	n, m := len(text), len(suffix)
	if m > n {
		return false
	}

	return text[n-m:] == suffix
}
