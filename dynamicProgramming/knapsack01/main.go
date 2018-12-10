package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// In this algorithm the W is the total weight of the sack
// wt is the weight of each item at an index
// val is the value of the item at the index
// and the n is the number of items
func knapSackGreedyRecursive(W int, wt []int, val []int, n int) int {
	if W == 0 || n == 0 {
		return 0
	}

	if wt[n-1] > W {
		return knapSackGreedyRecursive(W, wt, val, n-1)
	}

	return max((val[n-1] + knapSackGreedyRecursive(W-wt[n-1], wt, val, n-1)), knapSackGreedyRecursive(W, wt, val, n-1))
}

func main() {
	fmt.Println("vim-go")

	fmt.Println(knapSackGreedyRecursive(50, []int{10, 20, 30}, []int{60, 100, 120}, 3))
}
