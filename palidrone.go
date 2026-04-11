package main

import (
	"fmt"
	"strconv"
)

func main() {
	isPalindrome(123)
}

//func isPalindrome(x int) bool {
//
//	if x < 0 || (x%10 == 0 && x != 0) {
//		return false
//	}
//
//	reversed := 0
//
//	for x > reversed {
//		reversed = (reversed * 10) + x%10
//		x = x / 10
//	}
//
//	fmt.Println(x == reversed || x == reversed/10)
//
//	return x == reversed || x == reversed/10
//
//}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x) // convert number to string
	i, j := 0, len(s)-1  // two pointers

	for i < j {
		if s[i] != s[j] {
			fmt.Println(false)
			return false // mismatch found, exit immediately
		}
		i++
		j--
	}
	return true
}
