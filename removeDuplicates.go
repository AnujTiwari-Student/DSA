package main

import "fmt"

func main() {
	removeDuplicates([]int{1, 2, 3, 3, 4})
	fmt.Println(removeDuplicates([]int{1, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {

	ptr1 := 0

	for ptr2 := 1; ptr2 < len(nums); ptr2++ {
		if nums[ptr2] != nums[ptr1] {
			ptr1++
			nums[ptr1] = nums[ptr2]
		}

	}

	return ptr1 + 1
}
