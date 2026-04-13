package main

import "fmt"

type NewNode struct {
	data int
	next *NewNode
}

func main() {
	head := &NewNode{data: 1}
	head.next = &NewNode{data: 1}
	head.next.next = &NewNode{data: 2}
	head.next.next.next = &NewNode{data: 2}
	head.next.next.next.next = &NewNode{data: 3}

	fmt.Println("Before:")
	printListII(head)

	head = removeDuplicateListII(head)

	fmt.Println("\nAfter:")
	printListII(head)
}

func removeDuplicateListII(head *NewNode) *NewNode {

	if head == nil {
		return nil
	}

	list := &NewNode{data: 0}
	dummy := list

	currentNode := head

	for currentNode != nil {
		if currentNode.next != nil && currentNode.data == currentNode.next.data {
			dupVal := currentNode.data
			for currentNode != nil && currentNode.data == dupVal {
				currentNode = currentNode.next
			}
			continue
		}
		dummy.next = currentNode
		dummy = dummy.next
		currentNode = currentNode.next
	}
	dummy.next = nil

	return list.next
}

func printListII(head *NewNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.data)
		head = head.next
	}
	fmt.Println("nil")
}
