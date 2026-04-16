package main

import "fmt"

type nn struct {
	Val  int
	Next *nn
}

func main() {
	l1 := &nn{Val: 2, Next: &nn{Val: 4, Next: &nn{Val: 3}}}

	l2 := &nn{Val: 5, Next: &nn{Val: 6, Next: &nn{Val: 4}}}

	fmt.Print("List 1: ")
	printListForAddingTwoNumber(l1)
	fmt.Print("List 2: ")
	printListForAddingTwoNumber(l2)

	result := addTwoNumbers(l1, l2)

	fmt.Print("Sum:    ")
	printListForAddingTwoNumber(result)
}

func printListForAddingTwoNumber(head *nn) {
	curr := head
	for curr != nil {
		fmt.Printf("%d", curr.Val)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.Next
	}
	fmt.Println()
}

func addTwoNumbers(l1 *nn, l2 *nn) *nn {
	dummy := &nn{Val: 0}
	list := dummy

	carry := 0
	sum := 0
	digitToStore := 0

	for l2 != nil || l1 != nil || carry != 0 {

		v1, v2 := 0, 0

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum = v1 + v2 + carry

		carry = sum / 10
		digitToStore = sum % 10

		list.Next = &nn{Val: digitToStore}
		list = list.Next

	}

	return dummy.Next
}
