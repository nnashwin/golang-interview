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

func knapsackGreedyDP(W int, wt []int, val []int, n int) int {
	var K [][]int
	for i := 0; i <= n; i++ {
		K = append(K, []int{})
		for w := 0; w <= W; w++ {
			K[i] = append(K[i], 0)
		}
	}

	for i := 0; i <= n; i++ {
		for w := 0; w <= W; w++ {
			if i == 0 || w == 0 {
				K[i][w] = 0
			} else if wt[i-1] <= w {
				K[i][w] = max(val[i-1]+K[i-1][w-wt[i-1]], K[i-1][w])
			} else {
				K[i][w] = K[i-1][w]
			}
		}
	}

	return K[n][W]
}

func main() {
	fmt.Println("vim-go")

	fmt.Println(knapSackGreedyRecursive(50, []int{10, 20, 30}, []int{60, 100, 120}, 3))
	fmt.Println(knapsackGreedyDP(50, []int{10, 20, 30}, []int{60, 100, 120}, 3))
}
