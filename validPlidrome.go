package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	result := isPalindromePhrase(s)

	fmt.Printf("result is %v\n", result)
}

func isPalindromePhrase(s string) bool {

	runes := []rune(s)
	i := 0
	j := len(s) - 1

	for i < j {

		if !isAlphanumeric(runes[i]) {
			i++
			continue
		}

		// 2. Move j backward if not alphanumeric
		if !isAlphanumeric(runes[j]) {
			j--
			continue
		}

		if unicode.ToLower(runes[i]) != unicode.ToLower(runes[j]) {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
