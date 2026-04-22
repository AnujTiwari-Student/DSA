package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3}
	result := containsNearByDuplicate(arr, 3)
	fmt.Println(result)
}

//func containsNearByDuplicate(nums []int, k int) bool {
//	seen := make(map[int]int)
//
//	for i, num := range nums {
//
//		if lastIndex, ok := seen[num]; ok && i-lastIndex <= k {
//			return true
//		}
//
//		seen[num] = i
//	}
//	return false
//}

func containsNearByDuplicate(arr []int, k int) bool {
	seen := make(map[int]struct{})

	for i, val := range arr {
		if _, ok := seen[val]; ok {
			return true
		}

		seen[val] = struct{}{}

		if i >= k {
			delete(seen, arr[i-k])
		}
	}
	return false
}
