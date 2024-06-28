package kata

// Gap finds the first pair of prime numbers (p1, p2) with a gap of 'g' between 'm' and 'n'.
func Gap(g, m, n int) []int {
	for m+g <= n {
		if IsPrime(m) {
			if IsPrime(m + g) {
				// Check for any primes between m and m + g
				isGapPrimeFree := true
				for i := m + 1; i < m+g; i++ {
					if IsPrime(i) {
						isGapPrimeFree = false
						break
					}
				}
				if isGapPrimeFree {
					return []int{m, m + g}
				}
			}
		}
		m++
	}
	return nil
}
