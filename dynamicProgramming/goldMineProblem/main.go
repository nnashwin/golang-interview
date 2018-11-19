package main

import "fmt"

func build2dSquare(length int) [][]int {
	var square [][]int
	for i := 0; i <= length; i++ {
		square = append(square, []int{})
		for j := 0; j <= length; j++ {
			square[i] = append(square[i], 0)
		}
	}

	return square
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// right = [row][col + 1]
// rightTop = [row - 1][col + 1]
// rightBottom = [row + 1][col + 1]

func getHighestGold(mine [][]int) [][]int {
	dp := build2dSquare(len(mine) - 1)
	for col := len(mine[0]) - 1; col >= 0; col-- {
		for row := 0; row < len(mine); row++ {
			right := 0
			if col < len(mine[0])-1 {
				right = dp[row][col+1]
			}

			rightTop := 0
			if row > 0 && col < len(mine[0])-1 {
				rightTop = dp[row-1][col+1]
			}

			rightBottom := 0

			if row < len(mine)-1 && col < len(mine[0])-1 {
				rightBottom = dp[row+1][col+1]
			}

			dp[row][col] = mine[row][col] + max(right, max(rightTop, rightBottom))
		}
	}
	return dp
}

func main() {
	mine := [][]int{[]int{1, 3, 3}, []int{2, 1, 4}, []int{0, 6, 4}}
	table := getHighestGold(mine)
	fmt.Println(table)
	fmt.Println(mine)
	res := table[0][0]
	for i := 1; i < len(table); i++ {
		res = max(res, table[i][0])
	}

	fmt.Println(res)
}
