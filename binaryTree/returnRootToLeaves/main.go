package main

import "fmt"

type BNode struct {
	Val         int
	Left, Right *BNode
}

func isLeafNode(node BNode) bool {
	return node.Right == nil && node.Left == nil
}

func preOrderTraverse(node *BNode, currPath []int, allPaths *[][]int) {
	if node == nil {
		return
	}
	currPath = append(currPath, node.Val)
	if isLeafNode(*node) == true {
		*allPaths = append(*allPaths, currPath)
	}
	if node.Left != nil {
		preOrderTraverse(node.Left, currPath, allPaths)
	}

	if node.Right != nil {
		preOrderTraverse(node.Right, currPath, allPaths)
	}
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
	fmt.Println("vim-go")
	// bst := createTree([]int{1, 2, 3, 4, 5})
	bst := newNode(1)
	bst.Left = newNode(2)
	bst.Right = newNode(3)
	bst.Right.Left = newNode(4)
	bst.Right.Right = newNode(5)
	fmt.Println(*bst)
	finalArr := [][]int{}
	preOrderTraverse(bst, []int{}, &finalArr)
	fmt.Println(finalArr)
}
