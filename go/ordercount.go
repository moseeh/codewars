package kata

// Tuple struct to hold character and its count
type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(s string) []Tuple {
	if s == "" {
		return []Tuple{}
	}

	charMap := make(map[rune]int)
	var order []rune

	for _, char := range s {
		if _, exists := charMap[char]; !exists {
			order = append(order, char)
		}
		charMap[char]++
	}

	var result []Tuple
	for _, char := range order {
		result = append(result, Tuple{Char: char, Count: charMap[char]})
	}

	return result
}
