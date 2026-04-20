package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	fmt.Println(numbers)

	run1 := twoSumII(numbers, 9)
	fmt.Println(run1)

	numbers2 := []int{2, 3, 4}
	fmt.Println(numbers2)

	run2 := twoSumII(numbers2, 6)
	fmt.Println(run2)

}

func twoSumII(numbers []int, target int) []int {

	l := 0
	r := len(numbers) - 1

	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l++
		} else {
			r--
		}

	}
	return []int{-1, -1}

}
