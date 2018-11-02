package main

import "fmt"

func min(a, b int) {
	if a < b {
		return a
	}

	return b
}

func paintHouses(costs [][]int) int {
	if len(costs) == 0 || len(costs[0]) == 0 {
		return 0
	}

	for i := 1; i < len(costs); i++ {
		costs[i][0] += min(costs[i-1][1], costs[i-1][2])
		costs[i][1] += min(costs[i-1][0], costs[i-1[2]])
		costs[i][2] += min(costs[i-1][0], costs[i-1][1])
	}

	m := len(costs) - 1

	return min(min(costs[m][0], costs[m][1]), costs[m][2])
}

func main() {
	fmt.Println("vim-go")
}
