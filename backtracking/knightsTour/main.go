package main

import (
	"errors"
	"fmt"
)

var xMoves = []int{2, 1, -1, -2, -2, -1, 1, 2}
var yMoves = []int{1, 2, 2, 1, -1, -2, -2, -1}

func isSafe(x, y int, sol [][]int) bool {
	return x >= 0 && x < len(sol) && y >= 0 && y < len(sol[0]) && sol[x][y] == -1
}

func recurFunc(x, y, moveCount int, sol [][]int, xMoves, yMoves []int, iter *int) bool {
	if moveCount == len(sol)*len(sol[0]) {
		return true
	}
	*iter = *iter + 1

	for k := 0; k < 8; k++ {
		nextX := x + xMoves[k]
		nextY := y + yMoves[k]
		// backtracking occurs here
		if isSafe(nextX, nextY, sol) {
			sol[nextX][nextY] = moveCount
			if recurFunc(nextX, nextY, moveCount+1, sol, xMoves, yMoves, iter) {
				return true
			} else {
				sol[nextX][nextY] = -1
			}
		}
	}

	return false

}

func findKnightsTour(boardSize int) ([][]int, error) {
	// first we initialize each location on the board to -1 to not visited
	chessboard := make([][]int, boardSize)
	for x := 0; x < boardSize; x++ {
		for y := 0; y < boardSize; y++ {
			chessboard[y] = append(chessboard[y], -1)
		}
	}

	chessboard[0][0] = 0
	iter := 0

	if !recurFunc(0, 0, 1, chessboard, xMoves, yMoves, &iter) {
		fmt.Printf("ran through %d iterations\n", iter)
		return nil, errors.New("The solution doesn't exist.  Failure")
	} else {
		fmt.Printf("ran through %d iterations\n", iter)
		return chessboard, nil
	}

	fmt.Printf("ran through %d iterations\n", iter)

	return chessboard, nil
}

func main() {
	fmt.Println(findKnightsTour(2))
	fmt.Println(findKnightsTour(3))
	fmt.Println(findKnightsTour(6))
	fmt.Println(findKnightsTour(7))
}
