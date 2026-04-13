package main

import "fmt"

type List struct {
	val  int
	next *List
}

type Head struct {
	head *List
}

func main() {
	head := &List{val: 1}
	head.next = &List{val: 2}
	head.next.next = &List{val: 2}
	head.next.next.next = &List{val: 2}
	head.next.next.next.next = &List{val: 3}

	fmt.Println("Before:")
	printList(head)

	head = removeListDuplicate(head)

	fmt.Println("\nAfter:")
	printList(head)
}

// func removeListDuplicate(head *List) *List {
//
//		if head == nil {
//			return head
//		}
//
//		list := &List{val: 0}
//		tail := list
//
//		cn := head
//
//		for cn != nil {
//			if cn.next == nil || cn.val != cn.next.val {
//				tail.next = cn
//				tail = tail.next
//			}
//			cn = cn.next
//		}
//		tail.next = nil
//
//		return list.next
//	}
func printList(head *List) {
	for head != nil {
		fmt.Printf("%d -> ", head.val)
		head = head.next
	}
	fmt.Println("nil")
}

func removeListDuplicate(head *List) *List {
	if head == nil {
		return head
	}

	cn := head

	for cn != nil && cn.next != nil {
		if cn.val == cn.next.val {
			cn.next = cn.next.next
		} else {
			cn = cn.next
		}
	}

	return head
}
