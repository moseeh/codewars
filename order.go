package kata

import "strings"

func Order(sentence string) string {
	if sentence == "" {
		return ""
	}

	words := strings.Fields(sentence)
	sortedWords := make([]string, len(words))

	for _, word := range words {
		for _, char := range word {
			if char >= '1' && char <= '9' {
				position := int(char - '1')
				if position < len(sortedWords) {
					sortedWords[position] = word
				}
				break
			}
		}
	}

	return strings.Join(sortedWords, " ")
}
