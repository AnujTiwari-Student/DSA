package main

import "fmt"

func main() {
	result := listInNumber([]int{2, 2, -4, 5, 6, 7}, -4)
	fmt.Printf("result is %v\n", result)
}

func listInNumber(numbers []int, x int) bool {

	if numbers == nil || len(numbers) == 0 {
		return false
	}

	for _, n := range numbers {
		if n == x {
			return true
		}
	}
	return false
}
