package main

import "fmt"

func main() {
	nums := []int{10, 2, -2, -20, 10}
	target := -10

	result := subArraySumNegatives(nums, target)
	fmt.Printf("Indices: %v\n", result)
}

func subArraySumNegatives(nums []int, target int) []int {
	prefixMap := make(map[int]int)
	prefixMap[0] = -1
	currentSum := 0

	for j, val := range nums {
		currentSum += val
		neededSum := currentSum - target

		if i, exists := prefixMap[neededSum]; exists {
			return []int{i + 2, j + 1}
		}

		prefixMap[currentSum] = j
	}

	return []int{-1}
}
