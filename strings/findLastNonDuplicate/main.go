package main

import "fmt"

func lastNonDuplicate(str string) (returnCh string) {
	intervals := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		intervals[str[i]]++
	}

	for i := len(str) - 1; i >= 0; i-- {
		if intervals[str[i]] == 1 {
			returnCh = string(str[i])
			return
		}
	}

	return
}

func main() {
	fmt.Println(lastNonDuplicate("cacb"))
	fmt.Println(lastNonDuplicate("bcbc"))
	fmt.Println(lastNonDuplicate("GeeksforGeeks"))
	fmt.Println(lastNonDuplicate("GeeksQuiz"))
}
