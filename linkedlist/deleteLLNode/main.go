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

func deleteListNode(entry Node, root Node) *Node {
	var indirect *Node
	indirect = &(root)
	for (*indirect) != entry {
		indirect = indirect.Next
	}
	*indirect = *entry.Next

	return indirect
}

func main() {
	list := createList([]int{1, 2, 3, 4, 5})
	list2 := &Node{6, list}
	for current := list2; current != nil; {
		fmt.Println(current)
		current = current.Next
	}

	list2 = deleteListNode(*list2, *list2)
	for current := list2; current != nil; {
		fmt.Println(current)
		current = current.Next
	}

	secList := newNode(1)
	third := newNode(2)
	fourth := newNode(3)
	secList.Next = third
	third.Next = fourth
	third = deleteListNode(*third, *secList)
	for current := secList; current != nil; {
		fmt.Println(current)
		current = current.Next
	}
}
