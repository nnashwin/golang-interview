package main

import (
	"fmt"
	"strings"
)

func filterAdd(add string) string {
	addArr := strings.Split(add, "")
	for i := len(addArr) - 1; i >= 0; i-- {
		if addArr[i] == "." {
			addArr = append(addArr[:i], addArr[i+1:]...)
		}

		if addArr[i] == "+" {
			addArr = addArr[:i]
		}
	}

	return strings.Join(addArr, "")
}

func numUniqueEmails(emails []string) int {
	strHash := make(map[string]int)
	count := 0
	for i := 0; i < len(emails); i++ {
		emailArr := strings.Split(emails[i], "@")
		add := filterAdd(emailArr[0])
		dom := emailArr[1]
		filteredEmail := add + "@" + dom
		if _, ok := strHash[filteredEmail]; ok == false {
			strHash[filteredEmail] = 1
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}
