package main

import "fmt"

type BNode struct {
	Val         int
	Left, Right *BNode
}

func spop(stack *[]BNode) *BNode {
	if len(*stack) == 0 {
		return nil
	}

	slength := len(*stack) - 1

	sVal := *stack
	// get value of last index from arr
	nodeToReturn := sVal[slength]

	// pop value off of array
	sVal = append(sVal[:slength], sVal[slength+1:]...)
	*stack = sVal

	return &nodeToReturn
}

func spush(stack *[]BNode, node BNode) {
	if node == (BNode{}) {
		return
	}

	*stack = append(*stack, node)
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
	} else {
		root.Right = bstInsert(root.Right, key)
	}

	return root
}

func findKBiggestNode(root *BNode, k int) *BNode {
	if root == nil {
		return nil
	}

	stack := []BNode{*root}
	currentNode := root
	for currentNode != nil || len(stack) > 0 {
		fmt.Println(currentNode)
		for currentNode.Right != nil {
			spush(&stack, *currentNode.Right)
			*currentNode = *currentNode.Right
		}

		currentNode = spop(&stack)
		k--
		if k == 0 {
			return currentNode
		} else if currentNode.Left != nil {
			spush(&stack, *currentNode.Left)
			*currentNode = *currentNode.Left
		}
	}

	return nil
}

func createBst(arr []int) *BNode {
	var root *BNode
	keys := []int{15, 10, 20, 8, 5, 2}
	for _, k := range keys {
		root = bstInsert(root, k)
	}

	return root
}

func main() {

	bst3 := createBst([]int{15, 10, 20, 8, 5, 2})

	fmt.Println(findKBiggestNode(bst3, 6))
}
