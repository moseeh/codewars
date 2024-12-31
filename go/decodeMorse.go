package kata

import (
	"strings"
)

func DecodeMorse(morseCode string) string {
	words := strings.Split(morseCode, "   ")
	decoded := ""

	for _, word := range words {
		letters := strings.Split(word, " ")

		for _, letter := range letters {
			decoded += MORSE_CODE[letter]
		}

		decoded += " "
	}

	return strings.TrimSpace(decoded)
}
