package main

import "fmt"

func buildTable(n int) []int {
	var table []int
	for i := 0; i < n; i++ {
		table = append(table, 0)
	}

	return table
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minimumStepsToOne(n int) int {
	dp := buildTable(n + 1)
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = 1 + dp[i-1]
		if i%2 == 0 {
			dp[i] = min(dp[i], 1+dp[i/2])
		}

		if i%3 == 0 {
			dp[i] = min(dp[i], 1+dp[i/3])
		}
	}

	return dp[n]
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(minimumStepsToOne(1))
	fmt.Println(minimumStepsToOne(2))
	fmt.Println(minimumStepsToOne(4))
	fmt.Println(minimumStepsToOne(7))
}
