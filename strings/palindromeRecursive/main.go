package main

import "fmt"

func isPalindromeRecursive(text string) bool {
	if len(text) <= 1 {
		return true
	}

	if text[0] != text[len(text)-1] {

		return false
	}

	return isPalindromeRecursive(text[1 : len(text)-1])
}

func main() {
	fmt.Println(isPalindromeRecursive("mom"))
	fmt.Println(isPalindromeRecursive("dad"))
	fmt.Println(isPalindromeRecursive("racecar"))

	fmt.Println(isPalindromeRecursive("amaryllis sillyrama"))
}
