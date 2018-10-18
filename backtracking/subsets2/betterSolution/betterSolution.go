package main

import (
	"fmt"
	"sort"
)

func findSubsets(inputArr []int) [][]int {
	sort.Ints(inputArr)
	subsets := [][]int{}
	backtrack(inputArr, &subsets, 0, []int{})
	return subsets
}

func backtrack(inputArr []int, subsets *[][]int, index int, tempArr []int) {
	newTemp := make([]int, len(tempArr))
	copy(newTemp, tempArr)
	*subsets = append(*subsets, newTemp)

	for i := index; i < len(inputArr); i++ {
		if i > index && inputArr[i] == inputArr[i-1] {
			continue
		}

		tempArr = append(tempArr, inputArr[i])
		backtrack(inputArr, subsets, i+1, tempArr)
		tempArr = tempArr[1:]
	}
}

func main() {
	fmt.Println(findSubsets([]int{1, 2, 2}))
}
