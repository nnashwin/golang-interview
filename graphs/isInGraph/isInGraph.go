package main

import (
	"fmt"
)

func BuildLL(strs []string) *List {
	list := NewList()
	for idx, _ := range strs {
		fmt.Println(string(strs[idx]))
		el := Element{}
		el.Val = string(strs[idx])
		list.InsertAtEnd(&el)
	}
	return list
}

func main() {
	strs1 := []string{"t", "X", "T", "x", "y", "s", "P", "x"}
	// strs2 := []string{"b", "r", "e", "a", "k", "t", "h", "e", "s", "p", "e", "l", "l"}
	//
	ll := BuildLL(strs1)
	fmt.Printf("%+v", ll)
}
