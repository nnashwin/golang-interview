package main

import (
	"fmt"
	"sort"
)

func findSubsets(inputArr []int) [][]int {
	sort.Ints(inputArr)
	subsets := [][]int{}
	backtrack(inputArr, &subsets, 0, []int{}, make(map[int]bool))
	return subsets
}

func backtrack(inputArr []int, subsets *[][]int, index int, tempArr []int, containMap map[int]bool) {
	newTemp := make([]int, len(tempArr))
	copy(newTemp, tempArr)
	var tempSum int
	for i := 0; i < len(newTemp); i++ {
		tempSum += newTemp[i]
	}

	if _, ok := containMap[tempSum]; ok == false {
		*subsets = append(*subsets, newTemp)
		containMap[tempSum] = true
	}

	for i := index; i < len(inputArr); i++ {
		tempArr = append(tempArr, inputArr[i])
		backtrack(inputArr, subsets, i+1, tempArr, containMap)
		tempArr = tempArr[1:]
	}
}

func main() {
	fmt.Println(findSubsets([]int{1, 2, 2}))
}
