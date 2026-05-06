package main

import "fmt"

func main() {
	s := "Gopher Space"
	result := reverseString(s)
	fmt.Printf("Result is: %q \n", result)
}

func reverseString(s string) string {

	runes := []rune(s)

	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}
