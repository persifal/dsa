package gcd

func gcdIter(a, b int) int {
	for a > 0 && b > 0 {
		if a >= b {
			a %= b
		} else {
			b %= a
		}
	}

	return max(a, b)
}

func gcdRec(a, b int) int {
	if a == 0 || b == 0 {
		return max(a, b)
	}

	return gcdRec(b, a%b)
}
