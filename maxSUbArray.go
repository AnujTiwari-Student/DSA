package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-1, 2, 3, 4, 5, 6, 7, 8, 9, -10}
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
