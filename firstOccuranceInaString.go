package main

import (
	"fmt"
)

func main() {

	result := strStr("sadbutsad", "sad")

	result2 := strStr("mississippi", "issip")

	fmt.Println(result)

	fmt.Println(result2)

}

func strStr(haystack string, needle string) int {

	hLen := len(haystack)
	needleLen := len(needle)

	for i := 0; i <= hLen-needleLen; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}

	return -1

}

//func strStr(haystack string, needle string) int {
//	return strings.Index(haystack, needle)
//}
