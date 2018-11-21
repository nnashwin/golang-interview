package main

import "fmt"

func makeChange(coins []int, lenCoins int, n int) int {
	dp := []int{}
	for i := 0; i <= n; i++ {
		dp = append(dp, 0)
	}

	dp[0] = 1

	for i := 0; i < lenCoins; i++ {
		for j := coins[i]; j <= n; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	fmt.Println(dp)

	return dp[n]
}

func main() {
	fmt.Println(makeChange([]int{1, 2, 3}, 3, 4))
	fmt.Println(makeChange([]int{2, 5, 3, 6}, 4, 10))
}
