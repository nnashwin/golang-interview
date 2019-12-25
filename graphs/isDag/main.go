package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type VisitedStatus int

const (
	NotVisited VisitedStatus = iota
	JustVisited
	PermVisited
)

type GraphNode struct {
	Children   []string `json:"children"`
	HasVisited VisitedStatus
}

type Graph struct {
	GraphNodes map[string]*GraphNode `json:"nodes"`
}

func readGraphFromJson(fileLoc string) Graph {
	// read from the data dir in the graphs folder
	jsonFile, err := os.Open(fileLoc)

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var g Graph
	err = json.Unmarshal(byteValue, &g)

	if err != nil {
		panic(err)
	}

	return g
}

func isDag(g Graph) bool {
	isDag := true
	for k, _ := range g.GraphNodes {

		graphHasCycle := doesGraphHaveCycle(&g.GraphNodes, k)
		if graphHasCycle == true {
			isDag = false
			break
		}
	}

	return isDag
}

func doesGraphHaveCycle(nodes *map[string]*GraphNode, n string) bool {
	node := (*nodes)[n]
	fmt.Printf("str: %s, node: %v+\n", n, node)

	// if the node HasVisited is true, that means it has been visited before.  This means the graph has a cycle and is not a dag
	if node.HasVisited == PermVisited {
		return false
	}

	if node.HasVisited == JustVisited {
		return true
	}

	// mark node while recursing through all children
	node.HasVisited = JustVisited

	for _, nodeStr := range node.Children {
		if doesGraphHaveCycle(nodes, nodeStr) == true {
			return true
		}
	}

	node.HasVisited = PermVisited

	return false
}

func main() {
	dagG := readGraphFromJson("../data/dag/sharedDag.json")
	fmt.Println(isDag(dagG))

	a := GraphNode{Children: []string{"b", "c"}}
	b := GraphNode{Children: []string{"c"}}
	c := GraphNode{Children: []string{"d", "a"}}
	d := GraphNode{Children: []string{"a"}}
	gm := make(map[string]*GraphNode)
	gm["a"] = &a
	gm["b"] = &b
	gm["c"] = &c
	gm["d"] = &d
	nonDagG := Graph{GraphNodes: gm}
	fmt.Println(isDag(nonDagG))
}
