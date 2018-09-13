package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	seenHash := make(map[string]int)
	currCount := 0
	longestCount := 0
	for _, char := range s {
		strChar := string(char)
		if _, ok := seenHash[strChar]; ok == false {
			currCount += 1
			seenHash[strChar] = 1
		} else {
			seenHash = make(map[string]int)
			currCount = 1
			seenHash[strChar] = 1
		}

		if currCount > longestCount {
			longestCount = currCount
		}
	}

	return longestCount
}

func main() {
	fmt.Println("vim-go")
	lengthOfLongestSubstring("abcabcabb")
}
