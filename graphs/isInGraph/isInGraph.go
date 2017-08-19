package main

import (
	"fmt"
)

func BuildLL(strs []string) *List {
	list := NewList()
	for idx, _ := range strs {
		el := Element{}
		el.Val = string(strs[idx])
		list.InsertAtEnd(&el)
	}
	return list
}

func IsInGraph(str string, l *List) (foundStr bool) {
	strHash := make(map[string]int)
	fmt.Println(strHash)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Val)
	}
	return
}

func main() {
	strs1 := []string{"t", "X", "T", "x", "y", "s", "P", "x"}
	// strs2 := []string{"b", "r", "e", "a", "k", "t", "h", "e", "s", "p", "e", "l", "l"}
	ll := BuildLL(strs1)
	IsInGraph("tyler", ll)
}
