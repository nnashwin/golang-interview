package main

import "fmt"

type BNode struct {
	Val         int
	Left, Right *BNode
}

func newNode(key int) *BNode {
	node := BNode{Val: key}
	return &node
}

func bstInsert(root *BNode, key int) *BNode {
	if root == nil {
		return newNode(key)
	}

	if key < root.Val {
		root.Left = bstInsert(root.Left, key)
	} else if key > root.Val {
		root.Right = bstInsert(root.Right, key)
	}

	return root
}

func createBstFromArr(arr []int) *BNode {
	var root *BNode
	for _, k := range arr {
		root = bstInsert(root, k)
	}

	return root
}

func CountNodesRecursive(count int, root *BNode) int {
	count = 1
	if root.Left != nil {
		count += CountNodesRecursive(count, root.Left)
	}

	if root.Right != nil {
		count += CountNodesRecursive(count, root.Right)
	}

	return count
}

func main() {
	bst := createBstFromArr([]int{1, 2, 3, 4, 10, 13, 19, 15})
	count := CountNodesRecursive(0, bst)

	fmt.Println(count)
}
