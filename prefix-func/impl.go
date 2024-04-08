package main

func prefixFunc(s string) []int {
	pref := make([]int, len(s))
	// abbababb
	for i := 1; i < len(s); i++ {
		k := pref[i-1]
		for k > 0 && s[k] != s[i] {
			k = pref[k - 1]
		}

		if s[i] == s[k] {
			k++
		}

		pref[i] = k
	}

	return pref
}
