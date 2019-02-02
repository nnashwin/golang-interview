package main

import (
	"fmt"
	"strings"
)

func canBeShifted(a string, b string) bool {
	aArr := strings.Split(a, "")
	for i := 0; i < len(a); i++ {
		if strings.Join(aArr, "") == b {
			return true
		}
		aArr = append(aArr[1:], aArr[0])
	}

	return false
}

func main() {
	fmt.Println(canBeShifted("abc", "bca"))
	fmt.Println(canBeShifted("abc", "dog"))
}
