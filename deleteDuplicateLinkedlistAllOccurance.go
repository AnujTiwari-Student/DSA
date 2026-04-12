package main

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func main() {
	deleteDuplicateLinkedList(&LinkNode{1, &LinkNode{2, &LinkNode{3, &LinkNode{4, &LinkNode{4, &LinkNode{5, nil}}}}}})
}

func deleteDuplicateLinkedList(head *LinkNode) *LinkNode {
	list := &LinkNode{Val: 0}
	tail := list

	fmt.Printf("Tail value = %v\n", tail.Val)

	cn := head

	for cn != nil {
		if cn.Next != nil && cn.Val == cn.Next.Val {
			dupVal := cn.Val
			for cn != nil && cn.Val == dupVal {
				cn = cn.Next
			}
			continue
		}
		tail.Next = cn
		tail = tail.Next
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
