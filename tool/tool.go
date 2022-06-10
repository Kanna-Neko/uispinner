package tool

func Gcd(a, b int64) int64 {
	if a == 0 || b == 0 {
		return a | b
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
