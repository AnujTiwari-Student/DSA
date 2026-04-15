package main

import "fmt"

func main() {
	fmt.Println("First and Last index range are:", searchRange([]int{5, 7, 7, 8, 8, 10}, 8))

	result := searchRange([]int{5, 7, 8, 8, 8, 10}, 8)
	fmt.Printf("First and Last index range are %v\n", result)
}

func searchRange(nums []int, target int) []int {

	first := -1
	last := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			if first == -1 {
				first = i
				fmt.Printf("First index range is %v\n", first)
			}
			last = i
			fmt.Printf("Last index range is %v\n", last)
		}
	}

	return []int{first, last}
}
