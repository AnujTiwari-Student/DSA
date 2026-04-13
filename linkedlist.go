package main

import (
	"errors"
	"fmt"
)

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

	fmt.Println("\n--- After Adding Head (Data should be inserted) ---")
	list.InsertToStart(50)
	list.printList()

	fmt.Println("\n--- Deleting 90 ---")
	err := list.DeleteAnyNode(90)
	if err != nil {
		fmt.Printf("Error: %s\n", err) // This will print your "not found" message
	}
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

func (list *LinkedList) DeleteAnyNode(data int) error {
	if list.head == nil {
		return errors.New("list is empty")
	}

	if list.head.data == data {
		list.head = list.head.next
		return nil
	}

	for cn := list.head; cn.next != nil; {
		if cn.next.data == data {
			cn.next = cn.next.next
			return nil
		}
		cn = cn.next

	}

	return fmt.Errorf("node with data %d not found in the list", data)

}

func (list *LinkedList) DeleteAllOccurrences(data int) {
	for list.head != nil && list.head.data == data {
		list.head = list.head.next
	}

	if list.head == nil {
		return
	}

	cn := list.head
	for cn.next != nil {
		if cn.next.data == data {
			// BYPASS the node
			cn.next = cn.next.next
		} else {
			cn = cn.next
		}
	}
}

func (list *LinkedList) printList() {
	// start at the beginning
	cn := list.head

	for cn != nil {
		fmt.Printf("Current node address: %p\n Node Data: %d\n Address of the next node: %p\n", cn, cn.data, cn.next)
		cn = cn.next
	}
}
