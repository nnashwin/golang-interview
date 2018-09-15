package main

import (
	"fmt"
)

func getMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func lengthOfLongestSubstring(s string) int {
	seenHash := make(map[string]int)
	ans := 0
	for i, j := 0, 0; j < len(s); j++ {
		strChar := string(s[j])
		if _, ok := seenHash[strChar]; ok == true {
			i = getMax(seenHash[strChar], i)
		}

		ans = getMax(ans, j-i+1)
		seenHash[strChar] = j + 1
	}

	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcabb"))
}
