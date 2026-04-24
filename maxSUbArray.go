package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(arr)
	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	currSum := 0

	maxSum := math.MinInt

	for _, val := range nums {
		currSum += val
		maxSum = max(currSum, maxSum)
		if currSum < 0 {
			currSum = 0
		}
	}

	return maxSum

}
