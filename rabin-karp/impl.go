package main

var (
	prime        = 9923
	alphabetSize = 256
)

func rabinKarp(s, t string) []int {
	n := len(s)
	m := len(t)
	var thash int
	var shash int
	hornerHash := 1
	for i := 0; i < len(t); i++ {
		thash = (thash*alphabetSize + int(t[i])) % prime
		shash = (shash*alphabetSize + int(s[i])) % prime

		if i < m-1 {
			hornerHash *= alphabetSize
			hornerHash %= prime
		}
	}

	r := []int{}
	for i := 0; i <= n-m; i++ {
		if thash == shash && s[i:i+m] == t {
			r = append(r, i)
		}

		if i < n-m {
			shash = nextHash(shash, hornerHash, s[i], s[i+m])
		}
	}

	return r
}

func nextHash(prevHash int, hornerHash int, oldChar, newChar byte) int {
	nextHash := prevHash
	nextHash -= (int(oldChar) * hornerHash) % prime
	nextHash += prime
	nextHash *= alphabetSize
	nextHash += int(newChar)
	nextHash %= prime

	return nextHash
}
