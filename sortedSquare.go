package main

import "fmt"

func main() {
	arr := []int{-4, -1, 0, 3, 10}

	result := sortedSquares(arr)
	fmt.Println(result)

}

func sortedSquares(arr []int) []int {
	l := 0
	r := len(arr) - 1

	result := make([]int, len(arr))

	p := len(arr) - 1

	for l <= r {
		lsq := arr[l] * arr[l]
		fmt.Printf("%d * %d = %d\n", arr[l], arr[l], lsq)
		rsq := arr[r] * arr[r]
		fmt.Printf("%d * %d = %d\n", arr[r], arr[r], rsq)

		if lsq > rsq {
			result[p] = lsq
			p--
			l++
		} else {
			result[p] = rsq
			p--
			r--
		}

	}
	return result

}
