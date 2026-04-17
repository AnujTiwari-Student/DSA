package main

import (
	"fmt"
)

func main() {
	result := searchInsert([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Printf("result is %v\n", result)

	result1 := searchInsert([]int{5, 7, 10}, 4)
	fmt.Printf("result1 is %v\n", result1)
}

//func searchInsert(nums []int, target int) int {
//
//	return sort.SearchInts(nums, target)
//}

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
