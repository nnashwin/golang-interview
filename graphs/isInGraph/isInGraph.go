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
	foundStr = false

	if str == "" {
		return
	}

	strHash := make(map[string]int)
	for e := l.Front(); e != nil; e = e.Next() {
		strHash[e.Val] += 1
	}

	for _, char := range str {
		if strHash[string(char)] == 0 {
			return
		} else {
			strHash[string(char)] -= 1
		}
	}

	foundStr = true
	return
}

func main() {
	strs1 := []string{"t", "X", "T", "x", "y", "s", "P", "x"}
	ll := BuildLL(strs1)
	fmt.Println(IsInGraph("tyler", ll))
}
