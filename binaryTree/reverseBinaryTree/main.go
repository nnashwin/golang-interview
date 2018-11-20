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

func reverseBinaryTree(root *BNode) *BNode {
	if root == nil {
		return root
	}

	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp

	if root.Left != nil {
		reverseBinaryTree(root.Left)
	}

	if root.Right != nil {
		reverseBinaryTree(root.Right)
	}

	return root
}

func inOrder(root *BNode, sl *[]int) []int {
	if root == nil {
		return *sl
	}

	if root.Left != nil {
		inOrder(root.Left, sl)
	}

	*sl = append(*sl, root.Val)

	if root.Right != nil {
		inOrder(root.Right, sl)
	}

	return *sl
}

func main() {
	bst1 := createTree([]int{15, 10, 20, 8, 5, 2})
	fmt.Println(inOrder(bst1, &[]int{}))
	rbst := reverseBinaryTree(bst1)

	fmt.Println("reverse starts here")
	fmt.Println(inOrder(rbst, &[]int{}))
}
