package main

import "fmt"

func main() {
	s := "Gopher Space"
	result := reverseString(s)
	fmt.Printf("Result is: %q \n", result)

	d := "Gopher Space"
	result1 := reverseBytes([]byte(d))
	fmt.Printf("Result is: %q \n", result1)
}

func reverseString(s string) string {

	runes := []rune(s)

	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func reverseBytes(b []byte) []byte {
	m := len(b)

	for i := 0; i < m/2; i++ {
		b[i], b[m-1-i] = b[m-1-i], b[i]
	}

	return b

}
