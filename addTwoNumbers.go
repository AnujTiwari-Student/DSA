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

	for l2 != nil || l1 != nil || carry != 0 {

		sum := carry

		if l1 != nil {
			sum += l1.Val
			fmt.Printf("l1 sum is with carry=> %d\n", sum)
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			fmt.Printf("l1+l2 sum is =>%d\n", sum)
			l2 = l2.Next
		}

		carry = sum / 10
		fmt.Printf("Carry is: =>%d\n", carry)
		fmt.Printf("Sum after carry is =>%d\n", sum)

		list.Next = &nn{Val: sum % 10}
		fmt.Printf("List.Next after carry is =>%d\n", list.Next)
		list = list.Next
		fmt.Printf("List after carry is =>%d\n", list.Val)

	}

	return dummy.Next
}
