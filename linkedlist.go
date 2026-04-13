package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {
	var list LinkedList

	list.InsertToLast(32)
	list.InsertToLast(30)
	list.InsertToLast(48)
	list.InsertToLast(2)

	fmt.Println("--- Before Deletion ---")
	list.printList()

	list.DeleteNode()

	fmt.Println("\n--- After Deleting Head (32 should be gone) ---")
	list.printList()

	fmt.Println("\n--- After Adding Head (Data should be inserted) ---")
	list.InsertToStart(50)
	list.printList()

}

func (list *LinkedList) InsertToLast(data int) {
	nn := &Node{data, nil}
	if list.head == nil {
		list.head = nn

	} else {
		// Current node which are we on
		cn := list.head
		// For loop to find the last element of the list
		for cn.next != nil {
			cn = cn.next
		}
		cn.next = nn
	}

}

func (list *LinkedList) InsertToStart(data int) {

	nn := &Node{data, nil}

	nn.next = list.head
	list.head = nn

}

func (list *LinkedList) DeleteNode() *Node {

	if list.head == nil {
		return nil
	}
	list.head = list.head.next
	return list.head
}

func (list *LinkedList) printList() {
	// start at the beginning
	cn := list.head

	for cn != nil {
		fmt.Printf("Current node address: %p\n Node Data: %d\n Address of the next node: %p\n", cn, cn.data, cn.next)
		cn = cn.next
	}
}
