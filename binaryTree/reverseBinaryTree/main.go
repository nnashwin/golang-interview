package main

import "fmt"

type BNode struct {
	Val         int
	Left, Right *BNode
}

func newNode(val int) *BNode {
	return &BNode{Val: val}
}

func bstInsert(root *BNode, val int) *BNode {
	if root == nil {
		return newNode(val)
	}

	if val < root.Val {
		root.Left = bstInsert(root.Left, val)
	} else {
		root.Right = bstInsert(root.Right, val)
	}

	return root
}

func createTree(keys []int) *BNode {
	var root *BNode
	for _, k := range keys {
		root = bstInsert(root, k)
	}

	return root
}

func main() {
	bst1 := createTree([]int{15, 10, 20, 8, 5, 2})
	fmt.Println(bst1)
}
