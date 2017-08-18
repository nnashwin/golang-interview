package main

type Element struct {
	next *Element
	list *List
	Val  interface{}
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}

	return nil
}

type List struct {
	root Element
	len  int
}

func New() *List { return new(List).Init() }

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}

	return l.root.next
}

func (l *List) InsertAtEnd(e *Element) *Element {
	if l.Front() == nil {
		l.root.next = e
		e.list = l
		l.len++
		return e
	}

	node := l.Front()
	for node != nil {
		node = node.Next()
	}

	node.next = e
	e.list = l
	l.len++

	return e
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.len = 0
	return l
}
