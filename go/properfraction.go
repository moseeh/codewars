package kata

func ProperFractions(n int) int {
	if n == 1 {
		return 0
	}

	result := n
	p := 2
	for p*p <= n {
		if n%p == 0 {
			result -= result / p
			for n%p == 0 {
				n /= p
			}
		}
		p++
	}
	if n > 1 {
		result -= result / n
	}

	return result
}
