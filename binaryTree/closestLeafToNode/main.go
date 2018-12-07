package main

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

func findLeafDown(root *BNode, lev int, minDist *int) {
	if root == nil {
		return
	}

	// If the root wasn't null, and we have a value, set the minDist to be the current level value
	if root.Left == nil && root.Right == nil {
		if lev < *minDist {
			*minDist = lev
			return
		}
	}

	findLeafDown(root.Left, lev+1, minDist)
	findLeafDown(root.Right, lev+1, minDist)
}

func main() {
	btree := createTree([]int{6, 1, 3, 5, 20, 7, 9, 13})
}
