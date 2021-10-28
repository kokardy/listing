package listing

/*
functions for test
*/
var factorial = make(map[int]int)

func fac(n int) int {
	if len(factorial) < 2 {
		factorial[0], factorial[1] = 1, 1
	}
	if v, ok := factorial[n]; ok {
		return v
	}
	factorial[n] = n * fac(n-1)
	return factorial[n]
}

func cCount(n, m int) (result int) {
	switch {
	case n == m:
		result = 1
	case m == 1 || m == 0:
		result = n
	default:
		result = cCount(n-1, m) + cCount(n-1, m-1)
	}
	return result
}

func pCount(n, m int) (result int) {
	return fac(n) / fac(n-m)
}

func hCount(n, m int) (result int) {
	return cCount(n+m-1, m)
}
