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

// GcdEx is an extended Euclid algorithm
func GcdEx(a, b int) (int, int, int) {
	u0, v0 := 1, 0
	u1, v1 := 0, 1
	for b != 0 {
		q := a / b
		u0, u1 = u1, u0-q*u1
		v0, v1 = v1, v0-q*v1
		a, b = b, a%b
	}
	return a, u0, v0
}
