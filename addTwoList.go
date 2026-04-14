package main

import "fmt"

type node struct {
	val  int
	next *node
}

func main() {}

func addTwoList(l1 *node, l2 *node) *node {

	dummyList := &node{}
	fmt.Printf("Adding list to dummyList\n", dummyList)

	return dummyList
}
