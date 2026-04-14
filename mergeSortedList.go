package main

import "fmt"

type n struct {
	val  int
	next *n
}

func main() {
	// Test Case:
	// List 1: 1 -> 3 -> 5
	// List 2: 2 -> 4 -> 6
	l1 := &n{val: 1, next: &n{val: 3, next: &n{val: 5}}}
	//l2 := &n{val: 2}
	l2 := &n{val: 2, next: &n{val: 4, next: &n{val: 6}}}

	fmt.Println("List 1:")
	printNewList(l1)
	fmt.Println("List 2:")
	printNewList(l2)

	merged := mergeTwoLists(l1, l2)

	fmt.Println("\nMerged List:")
	printNewList(merged)
}

func printNewList(head *n) {
	for head != nil {
		fmt.Printf("%d -> ", head.val)
		head = head.next
	}
	fmt.Println("nil")
}

func mergeTwoLists(list1 *n, list2 *n) *n {
	dummy := &n{val: 0}
	nn := dummy

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	for list1 != nil && list2 != nil {
		if list1.val <= list2.val {
			nn.next = list1
			nn = nn.next
			list1 = list1.next
		} else {
			nn.next = list2
			nn = nn.next
			list2 = list2.next
		}
	}

	if list1 != nil {
		nn.next = list1
	} else {
		nn.next = list2
	}

	return dummy.next
}
