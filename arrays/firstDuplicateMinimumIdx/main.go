package main

import (
	"fmt"
	"math"
)

type CountHashMapper struct {
	Occurrences, SecondIdx int
}

func firstDuplicate(ints []int) int {
	// build a duplicate mechanism
	countHash := make(map[int]*CountHashMapper)
	for idx, val := range ints {
		if _, ok := countHash[val]; ok == false {
			countHash[val] = &CountHashMapper{
				Occurrences: 1,
				SecondIdx:   -1,
			}
		} else {
			(*countHash[val]).Occurrences += 1
			if countHash[val].Occurrences == 2 {
				(*countHash[val]).SecondIdx = idx
			}
		}
	}

	returnVal, duplicateIdx := -1, math.MaxInt32

	for k, v := range countHash {
		if v.Occurrences > 1 {
			if v.SecondIdx < duplicateIdx {
				returnVal = k
			}
		}
	}

	return returnVal
}

func main() {
	fmt.Println("vim-go")
	// test cases
	fmt.Println(firstDuplicate([]int{1, 2, 3, 4}) == -1)
	fmt.Println(firstDuplicate([]int{1, 2, 3, 4, 1}) == 1)
	fmt.Println(firstDuplicate([]int{1, 2, 3, 4, 2, 1}) == 2)
	fmt.Println(firstDuplicate([]int{6, 8, 1, 4, 2, 8, 6}) == 8)
}
