package kata

func FindOdd(seq []int) int {

	counts := make(map[int]int)
	for _, num := range seq {
		counts[num]++
	}
	var oddnum int
	for num, count := range counts {
		if count%2 != 0 {
			oddnum = num
			break
		}
	}
	return oddnum // your code here
}
