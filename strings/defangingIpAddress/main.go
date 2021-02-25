package main

import "fmt"

func defangIPaddr(address string) string {
	var ret string
	for _, val := range address {
		char := string(val)
		if string(char) == "." {
			ret += "[.]"
			continue
		}

		ret += char
	}

	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(defangIPaddr("255.100.50.0"))
}
