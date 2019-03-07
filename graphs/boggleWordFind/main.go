package main

import "fmt"

var boggleBoard = [3][3]string{{"A", "P", "P"},
	{"C", "A", "L"},
	{"T", "R", "E"}}

var dict = []string{"APPLE", "CAR", "CAT", "CARROT"}

func isWord(str string) bool {
	for i := 0; i < len(dict); i++ {
		if str == dict[i] {
			return true
		}
	}
	return false
}

func findWordsUtil(i, j int, visited *[][]bool, str *string) {
	(*visited)[i][j] = true
	*str = *str + boggleBoard[i][j]
	if isWord(*str) {
		fmt.Println(*str)
	}

	for row := i - 1; row <= i+1 && row < len(boggleBoard); row++ {
		for col := j - 1; col <= j+1 && col < len(boggleBoard[0]); col++ {
			if row >= 0 && col >= 0 && !(*visited)[row][col] {
				findWordsUtil(row, col, visited, str)
			}
		}
	}

	*str = (*str)[:len(*str)-1]
	(*visited)[i][j] = false
}

func findWords() {
	var visited [][]bool
	var str string
	for i := 0; i < 3; i++ {
		visited = append(visited, []bool{})
		for j := 0; j < 3; j++ {
			visited[i] = append(visited[i], false)
		}
	}

	for i := 0; i < len(boggleBoard); i++ {
		for j := 0; j < len(boggleBoard[0]); j++ {
			findWordsUtil(i, j, &visited, &str)
		}
	}
}

func main() {
	findWords()
}
