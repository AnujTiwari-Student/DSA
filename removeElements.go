package main

import "fmt"

func main() {
	result := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	fmt.Println(result)
}

func removeElement(nums []int, val int) int {

	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	nums = nums[:k]

	fmt.Println(nums)

	return len(nums)

}
