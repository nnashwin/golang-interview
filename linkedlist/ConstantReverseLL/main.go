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

func reverseList(root *Node) *Node {
	var prev *Node
	for current := root; current != nil; {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func main() {
	list := createList([]int{1, 2, 3, 4, 5})
	for current := list; current != nil; {
		fmt.Println(current)
		current = current.Next
	}
	reversed := reverseList(list)

	for current := reversed; current != nil; {
		fmt.Println(current)
		current = current.Next
	}
}
