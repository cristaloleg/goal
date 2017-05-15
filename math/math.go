package mathm

// Abs returns absolute value
func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

// Gcd returns greatest common divisor
func Gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// Lcm returns least common multiple
func Lcm(a, b int) int {
	return a * Gcd(a, b) / b
}
