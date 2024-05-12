package kata

func Tribonacci(signature [3]float64, n int) []float64 {
	if n == 0 {
		return []float64{}
	}

	if n <= 2 {
		return signature[:n]
	}
	s := make([]float64, n)
	s[0] = signature[0]
	s[1] = signature[1]
	s[2] = signature[2]

	for i := 3; i < n; i++ {
		s[i] = s[i-1] + s[i-2] + s[i-3]
	}
	return s
}
