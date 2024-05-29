package kata

import (
	"strconv"
	"strings"
)

func EncryptThis(text string) string {
	if text == "" {
		return ""
	}
	textarr := strings.Fields(text)
	for j, word := range textarr {
		s := strconv.Itoa(int(word[0]))
		if len(word) > 3 {
			s += string(word[len(word)-1]) + word[2:len(word)-1] + string(word[1])
		} else if len(word) == 3 {
			s += string(word[2]) + string(word[1])
		} else if len(word) == 2 {
			s += string(word[1])
		}
		textarr[j] = s
	}
	return strings.Join(textarr, " ")
}
