package kata

func Persistence(n int) int {
	count := 0
	num := n

	for num >= 10 {
		product := 1
		for num > 0 {
			product *= num % 10
			num /= 10
		}
		num = product
		count++
	}

	return count
}
