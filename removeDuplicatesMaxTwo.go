package main

import "fmt"

func main() {
	fmt.Println(removeDuplicatesMaxTwo([]int{1, 2, 2, 3, 4, 5}))
}

func removeDuplicatesMaxTwo(nums []int) int {
	ptr1 := 0

	var maps = make(map[int]int)
	maps[nums[0]] = 1

	for ptr2 := 1; ptr2 < len(nums); ptr2++ {
		if nums[ptr2] != nums[ptr1] {
			ptr1++
			nums[ptr1] = nums[ptr2]
		} else if maps[nums[ptr1]] < 2 {
			ptr1++
			nums[ptr1] = nums[ptr2]
		}
		maps[nums[ptr2]]++
	}
	return ptr1 + 1
}
