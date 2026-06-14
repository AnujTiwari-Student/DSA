package main

import "fmt"

func main() {
	arr := []int{9, 9}
	result := plusOne(arr)
	fmt.Printf("%#v\n", result)
	// Output: []int{1, 0, 0}
}

func plusOne(digits []int) []int {
	lastIndex := len(digits) - 1

	for i := lastIndex; i >= 0; i-- {
		digits[i]++
		if digits[i] > 9 {
			digits[i] = 0
		} else {
			return digits
		}
	}

	digits = append(digits, 0)

	digits[0] = 1

	return digits

}
