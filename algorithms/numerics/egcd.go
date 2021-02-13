package numerics

/// EXTENDED EUCLIDIAN ALGORITHM
// Author: kry127

// returns x, y: a*x + b*y = gcd(a, b)
func Egcd(a, b uint64) (int64, int64) {
	if a < b {
		return Egcd(b, a)
	}
	if b == 0 {
		return 1, 0
	}

	c := a % b
	r1, r2 := Egcd(b, c)
	// r1 * b + r2 * c = gcd(b, c)
	// r1 * b + r2 * (a % b) = gcd(a, b)
	// r1 * b + r2 * [(a % b) + a/b * b - a/b * b] = gcd(a, b)
	// r1 * b + r2 * [a - a/b * b] = gcd(a, b)
	// r2 * a + r1 * b - r2 * a/b * b = gcd(a, b)
	// r2 * a + [r1 - r2 * a/b] * b = gcd(a, b)

	return r2, r1 - r2 * int64(a/b)
}
