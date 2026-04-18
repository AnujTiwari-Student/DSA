package main

import "fmt"

func main() {

	//	arr = []int{1,2,3,7,5}, Target = 12
	//	we have to return [i+1, j+1]
	//	in this case its 2 and 4 :=> output [2, 4]
	//	First index of sub array in the example will be 1 and last 3 so output => [ 1+1 , 3+1 ]

	result := indexedOfSubArraySum([]int{1, 2, 3, 7, 5}, 12)
	fmt.Printf("%d\n", result)

}

func indexedOfSubArraySum(array []int, k int) []int {

	sum := 0
	i := 0

	for j := 0; j < len(array); j++ {
		sum += array[j]

		for sum > k && i < j {
			sum -= array[i]
			i++
			fmt.Printf("i is on: %d\n", i)
		}

		if sum == k {
			return []int{i + 1, j + 1}
		}

		fmt.Printf("J is on index: %d\n", j)

		fmt.Printf("Sum: %d\n", sum)

	}

	return []int{-1}
}
