package main

func zFunc(s string) []int {
	n := len(s)
	zf := make([]int, n)
	l, r := 0, 0
	// sad#sadbutsad
	for i := 1; i < n; i++ {
		//
		if i <= r {
			zf[i] = min(r-i+1, zf[i-l])
		}

		for zf[i]+i < n && s[zf[i]] == s[zf[i]+i] {
			zf[i]++
		}

		// match ends after 'r' (right bound)
		if zf[i]-1+i > r {
			r = zf[i] - 1 + i
			l = i
		}
	}

	return zf
}
