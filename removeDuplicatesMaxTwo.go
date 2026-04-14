package main

import "fmt"

func main() {
	fmt.Println(removeDuplicatesMaxTwo([]int{1, 1, 2, 2, 2, 3, 4, 5, 50, 100}))
}

func removeDuplicatesMaxTwo(nums []int) int {
	ptr1 := 2
	if len(nums) <= 2 {
		return len(nums)
	}
	for ptr2 := 2; ptr2 < len(nums); ptr2++ {
		if nums[ptr2] != nums[ptr1-2] {
			nums[ptr1] = nums[ptr2]
			ptr1++
		}
	}
	return ptr1
}
