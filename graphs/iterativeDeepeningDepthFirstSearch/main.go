package main

import (
	"fmt"
	gu "github.com/tlboright/golang-interview/graphs/utils"
)

func main() {
	sharedDagG := gu.ReadGraphFromJson("../data/dag/sharedDag.json")

	fmt.Println(sharedDagG)
}
