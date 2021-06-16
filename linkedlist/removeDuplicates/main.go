package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// (1) -> (1) -> (1)
// (1) -> (1) -> nil

func deleteDuplicates(head *ListNode) *ListNode {
	ret := head
	for head.Next != nil {
		if head.Val == (head.Next).Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	return ret.Next
}

func main() {
	nodes := []int{1, 1, 1, 3, 4, 5, 5, 5, 5, 5}
	head := &ListNode{}
	curr := head
	for i := 0; i < len(nodes); i++ {
		node := ListNode{Val: nodes[i], Next: nil}
		head.Next = &node
		head = head.Next
	}

	curr = deleteDuplicates(curr)

	for curr != nil {
		fmt.Println("nodeVal: ", curr.Val)
		curr = curr.Next
	}

}
