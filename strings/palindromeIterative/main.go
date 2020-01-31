package main

import "fmt"

func isPalindromeIterative(text string) bool {
	for i, l := 0, len(text)-1; i < len(text)/2; i++ {
		if text[i] != text[l-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindromeIterative("mom"))
	fmt.Println(isPalindromeIterative("dad"))
	fmt.Println(isPalindromeIterative("maam"))
	fmt.Println(isPalindromeIterative("racecar"))
	fmt.Println(isPalindromeIterative("amaryllis sillyrama"))
}
