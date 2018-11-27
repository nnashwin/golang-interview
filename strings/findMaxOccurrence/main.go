package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func firstMaxOccurrence(str string) (returnCh string) {
	intervals := make(map[rune]int)
	maxVal := 0

	for _, c := range str {
		intervals[c]++
		maxVal = max(maxVal, intervals[c])
	}

	for _, c := range str {
		if intervals[c] == maxVal {
			returnCh = string(c)
			return
		}
	}

	return
}

func main() {
	fmt.Println(firstMaxOccurrence("cacb"))
	fmt.Println(firstMaxOccurrence("bcbc"))
	fmt.Println(firstMaxOccurrence("GeeksforGeeks"))
	fmt.Println(firstMaxOccurrence("GeeksQuiz"))
}
