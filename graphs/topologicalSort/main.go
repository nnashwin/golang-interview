package main

import (
	"fmt"
	gu "github.com/tlboright/golang-interview/graphs/utils"
)

func topologicalSort(g gu.Graph) (sortedNodes []string) {
	for k, _ := range g.GraphNodes {
		visit(&g.GraphNodes, k, &sortedNodes)
	}

	return
}

func visit(nodes *map[string]*gu.GraphNode, n string, sortedNodes *[]string) {
	node := (*nodes)[n]

	if node.HasVisited == gu.PermVisited {
		return
	}

	if node.HasVisited == gu.JustVisited {
		panic("The graph you are sorting is not a dag.  Topological Sort can only be used to sort Dags.")
	}

	// mark node while its dependencies are being evaluated
	node.HasVisited = gu.JustVisited

	for _, nodeStr := range node.Children {
		visit(nodes, nodeStr, sortedNodes)
	}

	node.HasVisited = gu.PermVisited
	*sortedNodes = append([]string{n}, *sortedNodes...)
}

func main() {
	sharedDagG := gu.ReadGraphFromJson("../data/dag/sharedDag.json")

	fmt.Println(topologicalSort(sharedDagG))

	readmeGraph := gu.ReadGraphFromJson("nodeDag.json")
	fmt.Println(topologicalSort(readmeGraph))
}
