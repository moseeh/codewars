package kata

func Thirt(n int) int {
	sequence := []int{1, 10, 9, 12, 3, 4}
	digits := []int{}
	n1 := n
	for n1 > 0 {
		digits = append(digits, n1%10)
		n1 /= 10
	}
	sum := 0
	for i := 0; i < len(digits); i++ {
		sum += digits[i] * sequence[i%len(sequence)]
	}
	if sum == n {
		return sum
	} else {
		return Thirt(sum)
	}
}
