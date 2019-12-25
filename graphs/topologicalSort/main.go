package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// declare an enum VisitedStatus to map out 3 states in the sort:
// NotVisited -- This GraphNode has never been seen before
// JustVisited -- The GraphNode was seen in the current DFS visit recursive stack.  This enables us to check for cycles
// HasVisited -- This node is already in the list of sorted nodes for the graph and should not be visited again
type VisitedStatus int

const (
	NotVisited VisitedStatus = iota
	JustVisited
	PermVisited
)

type GraphNode struct {
	Val        int64    `json:"val"`
	Children   []string `json:"children"`
	HasVisited VisitedStatus
}

type Graph struct {
	GraphNodes map[string]*GraphNode `json:"nodes"`
}

func topologicalSort(g Graph) (sortedNodes []string) {
	for k, _ := range g.GraphNodes {
		visit(&g.GraphNodes, k, &sortedNodes)
	}

	return
}

func visit(nodes *map[string]*GraphNode, n string, sortedNodes *[]string) {
	node := (*nodes)[n]

	if node.HasVisited == PermVisited {
		return
	}

	if node.HasVisited == JustVisited {
		panic("The graph you are sorting is not a dag.  Topological Sort can only be used to sort Dags.")
	}

	// mark node while its dependencies are being evaluated
	node.HasVisited = JustVisited

	for _, nodeStr := range node.Children {
		visit(nodes, nodeStr, sortedNodes)
	}

	node.HasVisited = PermVisited
	*sortedNodes = append([]string{n}, *sortedNodes...)
}

func main() {
	// read from the data dir in the graphs folder
	jsonFile, err := os.Open("../data/dag/sharedDag.json")

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

	topologicalSort(g)
}
