package main

import "fmt"

type un struct {
	Val  int
	Next *un
}

func main() {
	list := &un{Val: 10}
	list.Next = &un{Val: 20}
	list.Next.Next = &un{Val: 30}
	list.Next.Next.Next = &un{Val: 40}
	list.Next.Next.Next.Next = &un{Val: 50}

	n := 2

	fmt.Println("Before:")
	printListForRemoving(list)

	removeNthFromEnd(list, n)

	fmt.Println("After:")
	printListForRemoving(list)
}

func removeNthFromEnd(head *un, n int) *un {

	if head == nil {
		return nil
	}

	dummy := &un{Val: 0, Next: head}
	fast := dummy
	slow := dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next

}

func printListForRemoving(head *un) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}
