package kata

func IsPrime(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 || n == 3 {
		return true
	} else if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
