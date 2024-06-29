package kata

func HighestRank(nums []int) int {
	mymap := make(map[int]int)
	for _, v := range nums {
		mymap[v]++
	}
	var mostFrequentNum int
	maxFrequency := 0
	for num, freq := range mymap {
		if freq > maxFrequency || (freq == maxFrequency && num > mostFrequentNum) {
			mostFrequentNum = num
			maxFrequency = freq
		}
	}
	return mostFrequentNum
}
