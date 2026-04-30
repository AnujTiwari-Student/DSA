package main

import "fmt"

func main() {
	s := "aaaabcabdecc"

	result := lengthOfLongestSubstring(s)
	fmt.Println(result)

}

func lengthOfLongestSubstring(s string) int {

	lastSeen := [128]int{}

	for i := range lastSeen {
		lastSeen[i] = -1
	}

	l := 0
	maxLen := 0

	for i, val := range s {

		if lastSeen[val] >= l {
			l = lastSeen[val] + 1
		}

		lastSeen[val] = i

		fmt.Printf("Last Seen Value is: %d \n", lastSeen[val])

		currentLen := i - l + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}

	}

	return maxLen

}
