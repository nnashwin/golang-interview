package main

import "fmt"

type BNode struct {
	Left, Right *BNode
	Val         int
}

func newNode(key int) *BNode {
	return &BNode{Val: key}
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

func lowestCommonAncestor(root *BNode, q int, p int) *BNode {
	if root == nil || root.Val == q || root.Val == p {
		return root
	}

	left := lowestCommonAncestor(root.Left, q, p)
	right := lowestCommonAncestor(root.Right, q, p)

	// either return root or return left / right
	if left != nil && right != nil {
		return root
	} else {
		if left != nil {
			return left
		}

		return right
	}
}

func main() {
	mid := newNode(5)
	mid.Right = newNode(8)
	mid.Left = newNode(3)
	mid.Right.Left = newNode(7)
	mid.Right.Right = newNode(9)
	mid.Left.Left = newNode(2)
	mid.Left.Right = newNode(4)
	fmt.Println(lowestCommonAncestor(mid, 7, 9))
	fmt.Println(lowestCommonAncestor(mid, 2, 4))
}
