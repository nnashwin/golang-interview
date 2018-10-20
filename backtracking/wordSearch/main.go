package main

import "fmt"

var board = [][]byte{
	[]byte{"A"[0], "B"[0], "C"[0], "E"[0]},
	[]byte{"S"[0], "F"[0], "C"[0], "S"[0]},
	[]byte{"A"[0], "D"[0], "E"[0], "E"[0]},
}

func inBounds(val int, length int) bool {
	return val >= 0 && val < length
}

func backtrack(board [][]byte, word string, i int, j int, result *bool) {
	if len(word) > 1 {
		if inBounds(i, len(board)) && inBounds(j-1, len(board[0])) && board[i][j-1] == word[1] {
			backtrack(board, word[1:], i, j-1, result)
		}

		if inBounds(i, len(board)) && inBounds(j+1, len(board[0])) && board[i][j+1] == word[1] {
			backtrack(board, word[1:], i, j+1, result)
		}

		if inBounds(i-1, len(board)) && inBounds(j, len(board[0])) && board[i-1][j] == word[1] {
			backtrack(board, word[1:], i-1, j, result)
		}

		if inBounds(i+1, len(board)) && inBounds(j, len(board[0])) && board[i+1][j] == word[1] {
			backtrack(board, word[1:], i+1, j, result)
		}
	} else {
		*result = true
	}
}

func exists(board [][]byte, word string) bool {
	result := false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				backtrack(board, word, i, j, &result)
			}
		}
	}

	return result
}

func main() {
	fmt.Println(exists(board, "ABCED"))
	fmt.Println(exists(board, "ABA"))
	fmt.Println(exists(board, "ABCE"))
	fmt.Println(exists(board, "ASFC"))
	fmt.Println(exists(board, "ASFF"))
}
