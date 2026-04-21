package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4}
	result := containsDuplicate(arr)
	fmt.Println(result)
}

func containsDuplicate(nums []int) bool {

	if len(nums) <= 1 {
		return false
	}

	seen := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}

	return false
}
