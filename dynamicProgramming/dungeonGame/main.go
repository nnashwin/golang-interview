package main

import "fmt"

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func findPath(grid [][]int) int {
	M := len(grid)
	N := len(grid[0])
	// initialize everything with max int
	hp := [][]int{}
	for i := 0; i < M+1; i++ {
		hp = append(hp, []int{})
		for j := 0; j < N+1; j++ {
			hp[i] = append(hp[i], MaxInt)
		}
	}

	hp[M][N-1] = 1
	hp[M-1][N] = 1

	for i := M - 1; i >= 0; i-- {
		for j := N - 1; j >= 0; j-- {
			need := min(hp[i+1][j], hp[i][j+1]) - grid[i][j]
			hp[i][j] = need
			if need <= 0 {
				hp[i][j] = 1
			}
		}
	}
	return hp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	testCase := [][]int{[]int{-2, -3, 3}, []int{-5, -10, 1}, []int{10, 30, -5}}
	fmt.Println(findPath(testCase))
}
