package kata

import "strings"

func High(s string) string {
	sArr := strings.Fields(s)
	wordscore := 0
	wordScoreArr := []int{}
	for _, word := range sArr {
		for _, v := range word {
			wordscore += int(v-'a') + 1
		}
		wordScoreArr = append(wordScoreArr, wordscore)
		wordscore = 0
	}
	return sArr[indexOfMax(wordScoreArr)]
}
func indexOfMax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	maxIndex := 0
	maxValue := arr[0]

	for i, value := range arr {
		if value > maxValue {
			maxValue = value
			maxIndex = i
		}
	}

	return maxIndex
}
