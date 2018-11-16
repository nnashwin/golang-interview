package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minNum := prices[0]
	maxDiff := 0
	for i := 1; i < len(prices); i++ {
		potentialDiff := prices[i] - minNum
		maxDiff = max(potentialDiff, maxDiff)

		minNum = min(minNum, prices[i])
	}

	return maxDiff

}

func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
