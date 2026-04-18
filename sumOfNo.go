package main

import "fmt"

func main() {
	result := sumOfNumbers(nil)
	fmt.Println(result)
}

func sumOfNumbers(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sum := 0

	//for i := 0; i < len(nums); i++ {
	//	sum += nums[i]
	//}

	for _, val := range nums {
		sum += val
	}

	return sum
}
