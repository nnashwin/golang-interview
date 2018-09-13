package main

import (
	"fmt"
)

func findMax(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

func lengthOfLongestSubstring(s string) int {
	seenHash := make(map[string]int)
	strLen := len(s)
	ans, i, j := 0, 0, 0
	for i < strLen && j < strLen {
		headStr := string(s[i])
		tailStr := string(s[j])
		if _, ok := seenHash[tailStr]; ok == false {
			seenHash[tailStr] = 1
			j += 1
			ans = findMax(ans, j-i)
		} else {
			delete(seenHash, headStr)
			i += 1
		}
	}

	return ans
}

func main() {
	fmt.Println("vim-go")
	lengthOfLongestSubstring("abcabcabb")
}
