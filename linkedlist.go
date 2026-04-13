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
	size int
}

func main() {
	var list LinkedList

	list.InsertToLast(32)
	list.InsertToLast(30)
	list.InsertToLast(48)
	list.InsertToLast(2)

	fmt.Println("--- Before Deletion ---")
	list.printList()

	fmt.Println("\n--- After Adding Head ---")
	list.InsertToStart(50)
	list.printList()

	fmt.Println("\n--- Deleting 90 ---")
	err := list.DeleteAnyNode(90)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	list.printList()

	fmt.Println("\n--- Adding 90 at position 6 ---")
	newErr := list.AddNode(90, 6)
	if newErr != nil {
		fmt.Printf("Error: %s\n", newErr)
	}
	list.printList()
}

func (list *LinkedList) InsertToLast(data int) {
	nn := &Node{data, nil}
	list.size++
	if list.head == nil {
		list.head = nn
		return
	}
	cn := list.head
	for cn.next != nil {
		cn = cn.next
	}
	cn.next = nn
}

func (list *LinkedList) InsertToStart(data int) {
	nn := &Node{data, nil}
	list.size++
	nn.next = list.head
	list.head = nn
}

func (list *LinkedList) DeleteAnyNode(data int) error {
	if list.head == nil {
		return errors.New("list is empty")
	}

	if list.head.data == data {
		list.head = list.head.next
		list.size--
		return nil
	}

	for cn := list.head; cn.next != nil; {
		if cn.next.data == data {
			cn.next = cn.next.next
			list.size--
			return nil
		}
		cn = cn.next
	}

	return fmt.Errorf("node with data %d not found", data)
}

func (list *LinkedList) DeleteAllOccurrences(data int) {
	for list.head != nil && list.head.data == data {
		list.head = list.head.next
		list.size--
	}

	if list.head == nil {
		return
	}

	for cn := list.head; cn.next != nil; {
		if cn.next.data == data {
			cn.next = cn.next.next
			list.size--
		} else {
			cn = cn.next
		}
	}
}

func (list *LinkedList) AddNode(data int, pos int) error {
	if pos < 1 || pos > list.size+1 {
		return fmt.Errorf("invalid position %d. Enter between 1 and %d", pos, list.size+1)
	}

	if pos == 1 {
		list.InsertToStart(data)
		return nil
	}

	nn := &Node{data, nil}
	cn := list.head
	count := 1

	for count < pos-1 {
		cn = cn.next
		count++
	}

	nn.next = cn.next
	cn.next = nn
	list.size++
	return nil
}

func (list *LinkedList) printList() {
	cn := list.head
	fmt.Printf("List Size: %d\n", list.size)
	for cn != nil {
		fmt.Printf("Addr: %p | Data: %d | Next: %p\n", cn, cn.data, cn.next)
		cn = cn.next
	}
}
