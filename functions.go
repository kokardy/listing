package listing

var factorial = make(map[int]int)

func Fac(n int) int {
	if len(factorial) < 2 {
		factorial[0], factorial[1] = 1, 1
	}
	if v, ok := factorial[n]; ok {
		return v
	}
	factorial[n] = n * Fac(n-1)
	return factorial[n]
}

func C(n, m int) (result int) {
	switch {
	case n == m:
		result = 1
	case m == 1 || m == 0:
		result = n
	default:
		result = C(n-1, m) + C(n-1, m-1)
	}
	return result
}

func P(n, m int) (result int) {
	return Fac(n) / Fac(n-m)
}

func H(n, m int) (result int) {
	return C(n+m-1, m)
}
