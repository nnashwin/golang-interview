package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func newNode(num int) *Node {
	return &Node{Data: num}
}

func createList(nums []int) *Node {
	root := newNode(nums[0])
	idx := root
	for i := 1; i <= len(nums)-1; i++ {
		num := newNode(nums[i])
		idx.Next = num
		idx = num
	}

	return root
}

func insertListNode(head **Node, val int) {
	for *head != nil {
		head = &((*head).Next)
	}
	*head = &Node{Data: val}
}

func main() {
	list := createList([]int{1, 2, 3, 4, 5})
	for current := list; current != nil; {
		fmt.Println(current)
		current = current.Next
	}
	insertListNode(&list, 6)

	for current := list; current != nil; {
		fmt.Println(current)
		current = current.Next
	}

	var headNode *Node
	insertListNode(&headNode, 5)
	fmt.Println(headNode)
}
