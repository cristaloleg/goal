package mathm

// Abs X
func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

// Gcd X
func Gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// Lcm X
func Lcm(a, b int) int {
	return a * Gcd(a, b) / b
}
