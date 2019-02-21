package main

import "fmt"

type MinNode struct {
	Val  int
	Min  int
	Prev *MinNode
}

type MinStack struct {
	Head *MinNode
	Min  int
}

func NewMinStack() *MinStack {
	ms := MinStack{}
	return &ms
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func (ms *MinStack) GetMin() int {
	return ms.Min
}

func (ms *MinStack) Pop() MinNode {
	var pn MinNode
	pn = *ms.Head

	ms.Head = ms.Head.Prev
	ms.Min = ms.Head.Min

	return pn
}

func (ms *MinStack) Push(val int) {
	if ms.Head == nil {
		ms.Head = &MinNode{Val: val, Min: val}
		ms.Min = val
		return
	}

	mn := MinNode{Val: val}
	mn.Min = Min(val, ms.Min)
	ms.Min = Min(val, ms.Min)

	mn.Prev = ms.Head
	ms.Head = &mn

	return
}

func (ms *MinStack) Top() MinNode {
	return *ms.Head
}

func main() {
	fmt.Println("vim-go")
	minStack := NewMinStack()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack)
	fmt.Println(*minStack.Head)
	fmt.Printf("Get Min %d\n", minStack.GetMin())
	fmt.Println(minStack.Pop())
	fmt.Printf("Get Top %v\n", minStack.Top())
	fmt.Printf("Get Min %d\n", minStack.GetMin())
}
