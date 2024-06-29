package main

import "fmt"

func main() {
	fmt.Println(wave("moses is a fool . "))

}
func wave(words string) []string {
	if words == "" {
		return []string{}
	}
	wordsrune := []rune(words)
	wordsarr := []string{}
	for i, v := range wordsrune {
		ws := wordsrune
		if v >= 'a' && v <= 'z' {
			ws[i] = v - 32
			wordsarr = append(wordsarr, string(wordsrune))
			ws[i] = v
		}
	}
	return wordsarr

}
