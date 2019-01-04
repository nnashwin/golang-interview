package main

import "fmt"

type BNode struct {
	Val         int
	Left, Right *BNode
}

func newNode(val int) *BNode {
	return &BNode{Val: val}
}

func bstInsert(val int, root *BNode) *BNode {
	if root == nil {
		return newNode(val)
	}

	if val < root.Val {
		root.Left = bstInsert(val, root.Left)
	} else {
		root.Right = bstInsert(val, root.Right)
	}

	return root
}

func createTree(keys []int) *BNode {
	var root *BNode
	for _, k := range keys {
		root = bstInsert(k, root)
	}

	return root
}

func main() {
	bst1 := createTree([]int{15, 20, 20, 8, 5, 2})
	fmt.Println(bst1)
}
