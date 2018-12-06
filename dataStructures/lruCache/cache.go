package main

import "fmt"

type CacheNode struct {
	Val, Loc   int
	Prev, Next *CacheNode
}

type LRUCache struct {
	Head, Tail *CacheNode
	NodeMap    map[int]*CacheNode
	Length     int
	Size       int
}

func newCache(size int) LRUCache {
	return LRUCache{Size: size, NodeMap: make(map[int]*CacheNode)}
}

func (cache *LRUCache) Put(loc int, val int) {
	node := &CacheNode{Val: val, Loc: loc}
	if cache.Head == nil {
		cache.Head = node
		cache.Tail = node
		cache.NodeMap[loc] = node
		cache.Length++
		return
	}

	// remove oldest element out of the cache
	if cache.Length == cache.Size {
		next := cache.Head.Next
		delete(cache.NodeMap, cache.Head.Loc)
		cache.Head = next
		cache.Length--
	}

	cache.NodeMap[loc] = node
	cache.Tail.Next = node
	cache.Tail = node
	cache.Length++
}

func (cache *LRUCache) Get(loc int) int {
	if _, ok := cache.NodeMap[loc]; ok == true {
		return cache.NodeMap[loc].Val
	}
	return -1
}

func (cache *LRUCache) Remove(loc int) {
	if _, ok := cache.NodeMap[loc]; ok == true {
		delete(cacheNodeMap, loc)
		cache.Length--
	}
}

func main() {
	fmt.Println("vim-go")
	c := newCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	fmt.Println(c.Get(1))
	c.Put(4, 4)
	fmt.Println(c)
	fmt.Println(c.Get(2))
	c.Put(1, 1)
}
