package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	list := &ListNode{Val: 0}
	tail := list

	fmt.Printf("Tail: %d\n", tail.Val)

	fmt.Println("List value at initialization is: ", list)

	cn := head

	fmt.Printf("Head: %d\n", cn.Val)

	for cn != nil {
		//fmt.Printf("head: %d\t Data: %d\n", *cn, cn.Val)

		if cn.Next == nil || cn.Val != cn.Next.Val {
			tail.Next = cn

			fmt.Printf("Assigning value to tail.next: %d\n", tail.Val)

			tail = tail.Next

			fmt.Printf("Assigning tail value of next tail: %d\n", tail.Val)
		}
		cn = cn.Next

	}
	tail.Next = nil

	result := list.Next
	for result != nil {
		fmt.Printf("%d -> ", result.Val)
		result = result.Next
	}
	fmt.Println("nil")

	return list.Next

}

func main() {
	deleteDuplicates(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}})
}
