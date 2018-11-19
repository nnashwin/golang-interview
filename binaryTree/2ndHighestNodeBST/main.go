package main

import "fmt"

type BNode struct {
	val         int
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

func findKBiggestNode(bst BNode) *BNode {
	if bst == (BNode{}) {
		return nil
	}
	return nil
}

func main() {
	fmt.Println("vim-go")
	bnode := BNode{val: 100}
	bnode2 := BNode{val: 101}
	bnode3 := BNode{val: 102}

	stack := []BNode{}

	spush(&stack, bnode)
	spush(&stack, bnode2)
	spush(&stack, bnode3)
	node := spop(&stack)
	fmt.Println(node)
	fmt.Println(stack)
}
