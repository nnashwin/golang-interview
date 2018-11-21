package main

import "fmt"

type BNode struct {
	Val         int
	Left, Right *BNode
}

func newNode(val int) *BNode {
	return &BNode{Val: val}
}

func isMirror(n1 *BNode, n2 *BNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1 != nil && n2 != nil && n1.Val == n2.Val {
		return isMirror(n1.Left, n2.Right) && isMirror(n2.Left, n1.Right)
	}

	return false
}

func isSymmetric(r *BNode) bool {
	return isMirror(r, r)
}

func main() {
	bst := newNode(1)
	bst.Left = newNode(2)
	bst.Right = newNode(2)
	bst.Left.Left = newNode(3)
	bst.Left.Right = newNode(4)
	bst.Right.Left = newNode(4)
	bst.Right.Right = newNode(3)

	bst2 := newNode(1)
	bst2.Left = newNode(2)
	bst2.Right = newNode(2)
	bst2.Left.Right = newNode(3)
	bst2.Right.Right = newNode(3)

	fmt.Println(isSymmetric(bst))
	fmt.Println(isSymmetric(bst2))
}
