package kata

func TwoSum(numbers []int, target int) [2]int {
	numarr := [2]int{}
	for i := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if target == numbers[i]+numbers[j] {
				numarr[0], numarr[1] = i, j
				return numarr
			}
		}
	}
	return numarr
}
