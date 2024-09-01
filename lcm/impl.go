package gcd

func gcd(a, b int) int {
	for a > 0 && b > 0 {
		if a >= b {
			a %= b
		} else {
			b %= a
		}
	}

	return max(a, b)
}
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
