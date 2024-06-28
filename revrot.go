package kata

import "strconv"

func chunck(s string) string {
	sum := 0
	for _, char := range s {
		digit, _ := strconv.Atoi(string(char))
		sum += digit
	}
	if sum%2 == 0 {
		// If sum of digits is divisible by 2, reverse the chunk
		return reverse(s)
	} else {
		// Otherwise, rotate the chunk to the left by one position
		return s[1:] + string(s[0])
	}
}

func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Revrot(str string, sz int) string {
	if sz <= 0 || str == "" || sz > len(str) {
		return ""
	}
	result := ""
	for i := 0; i+sz <= len(str); i += sz {
		chunk := str[i : i+sz]
		result += chunck(chunk)
	}
	return result
}
