package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{2, 3, 7, 8, 9}

	result := getCommon(nums1, nums2)
	fmt.Println(result)

}

func getCommon(nums1 []int, nums2 []int) int {

	p1 := 0
	p2 := 0

	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] < nums2[p2] {
			p1++
		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {
			return nums1[p1]
		}
	}

	return -1
}
