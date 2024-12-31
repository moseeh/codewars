package kata

func GetCount(str string) (count int) {
	count = 0
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, v := range str {
		for _, vowel := range vowels {
			if string(v) == vowel {
				count++
			}
		}
	}
	return count
}
