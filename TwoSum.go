package main

import "fmt"

func main() {
	twoSum([]int{3, 2, 4}, 6)
}

func twoSum(nums []int, target int) []int {
	var result = []int{0}

	var maps = make(map[int]int)

	for i, _ := range nums {

		complement := target - nums[i]

		if idx, ok := maps[complement]; ok && idx != i {

			fmt.Println("Index", idx)

			result = []int{idx, i}
			fmt.Println(result)
		}
		maps[nums[i]] = i
		fmt.Println(maps)

	}
	return result
}
