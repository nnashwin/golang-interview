package main

import "fmt"

func firstNonDuplicate(str string) (returnCh string) {
	intervals := make(map[rune]int)

	for _, c := range str {
		intervals[c]++
	}

	for _, c := range str {
		if intervals[c] == 1 {
			returnCh = string(c)
			return
		}
	}

	return
}

func main() {
	fmt.Println(firstNonDuplicate("cacb"))
	fmt.Println(firstNonDuplicate("bcbc"))
	fmt.Println(firstNonDuplicate("GeeksforGeeks"))
	fmt.Println(firstNonDuplicate("GeeksQuiz"))
}
