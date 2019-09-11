package main

import "fmt"

type LfuCache struct {
	ByKey    map[string]*FrequencyNode
	FreqHead *FrequencyNode
}

type FrequencyNode struct {
	Value      int
	Items      map[string]string
	Prev, Next *FrequencyNode
}

type LfuItem struct {
	Data   string
	Parent *FrequencyNode
}

func newLfuCache() LfuCache {
	var lc LfuCache
	lc.ByKey = make(map[string]*FrequencyNode)
	lc.FreqHead = &FrequencyNode{}

	return lc
}

func newFrequencyNode() FrequencyNode {
	var n FrequencyNode
	n.Value = 0
	n.Prev = &n
	n.Next = &n

	return n
}

func getNewFreqNode(value int, prev, next *FrequencyNode) FrequencyNode {
	nn := newFrequencyNode()
	nn.Value = value
	nn.Prev = prev
	nn.Next = next
	prev.Next = nn
	next.Prev = nn

	return nn
}

func newLfuItem(data string, parent *FrequencyNode) LfuItem {
	return LfuItem{data, parent}
}

func (c *LfuCache) Insert(key, val string) {
	if c.ByKey[key] {
		return nil
	}

	freq := c.FreqHead.Next
	if freq.Value != 1 {
		freq = getNewFreqNode(1, c.FreqHead, freq)
	}

	freq.Items.add(key, val)
	c.ByKey[key] = newLfuItem(val, freq)
}

func main() {

	lfuCache := newLfuCache()
	fmt.Println("%v", lfuCache.FreqHead)
}
