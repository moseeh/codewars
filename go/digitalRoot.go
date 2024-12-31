package kata

func DigitalRoot(n int) int {
	if n < 10 {
		return n
	} else {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		return DigitalRoot(sum)
	}
}
