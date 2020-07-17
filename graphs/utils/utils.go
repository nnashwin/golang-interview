package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Graph struct {
	GraphNodes map[string]*GraphNode `json:"nodes"`
}

type GraphNode struct {
	Children   []string `json:"children"`
	HasVisited VisitedStatus
}

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

func ReadGraphFromJson(fileLoc string) Graph {
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
