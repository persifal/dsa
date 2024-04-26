package main

func Kmp(t, p string) int {
	n := len(t)
	m := len(p)
	pref := prefix(p)
	i, j := 0, 0
	for /*remain length in text*/ (n - i) >= /*remain length in pattern*/ (m - j) {
		if t[i] == p[j] {
			i++
			j++
		}

		if j == m {
			return i - m
		} else {
			if t[i] != p[j] && i < n {
				if j != 0 {
					j = pref[j-1]
				} else {
					i++
				}
			}
		}
	}

	return -1
}

func prefix(s string) []int {
	pref := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		k := pref[i-1]
		for k > 0 && s[k] != s[i] {
			k = pref[k-1]
		}

		if s[i] == s[k] {
			k++
		}

		pref[i] = k
	}

	return pref
}
