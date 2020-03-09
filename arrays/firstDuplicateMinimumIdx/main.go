package main

import (
	"fmt"
	"math"
)

type CountHashMapper struct {
	Occurrences, SecondIdx int
}

func findFirstDuplicate(ints []int) int {
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

func findFirstDuplicateOptimized(ints []int) int {
	// build a duplicate mechanism
	countHash := make(map[int]*CountHashMapper)

	returnVal, duplicateIdx := -1, math.MaxInt32
	for idx, val := range ints {
		if _, ok := countHash[val]; ok == false {
			countHash[val] = &CountHashMapper{
				Occurrences: 1,
			}
		} else {
			(*countHash[val]).Occurrences += 1
			if countHash[val].Occurrences == 2 {
				if idx < duplicateIdx {
					duplicateIdx = idx
					returnVal = val
				}
			}
		}
	}

	return returnVal
}

func main() {
	// test cases
	fmt.Println(`
	    findFirstDuplicateUnoptimized
	`)
	fmt.Println(findFirstDuplicate([]int{1, 2, 3, 4}) == -1)
	fmt.Println(findFirstDuplicate([]int{1, 2, 3, 4, 1}) == 1)
	fmt.Println(findFirstDuplicate([]int{1, 2, 3, 4, 2, 1}) == 2)
	fmt.Println(findFirstDuplicate([]int{6, 8, 1, 4, 2, 8, 6}) == 8)

	fmt.Println(`
	    findFirstDuplicateOptimized
	`)
	fmt.Println(findFirstDuplicateOptimized([]int{1, 2, 3, 4}) == -1)
	fmt.Println(findFirstDuplicateOptimized([]int{1, 2, 3, 4, 1}) == 1)
	fmt.Println(findFirstDuplicateOptimized([]int{1, 2, 3, 4, 2, 1}) == 2)
	fmt.Println(findFirstDuplicateOptimized([]int{6, 8, 1, 4, 2, 8, 6}) == 8)
}
