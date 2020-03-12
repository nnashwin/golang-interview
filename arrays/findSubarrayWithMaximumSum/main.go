package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maximumSumInefficient(arr []int) int {
	globalMax := arr[0]
	for i := 0; i < len(arr); i++ {
		localMax := arr[i]
		for j := i + 1; j < len(arr); j++ {
			localMax = max(localMax+arr[j], arr[j])
			globalMax = max(localMax, globalMax)
		}
	}

	return globalMax
}

func maxSum(arr []int) int {
	currentMax, maxEndingHere := arr[0], arr[0]
	for idx := 1; idx < len(arr); idx++ {
		val := arr[idx]
		maxEndingHere = max(maxEndingHere+val, val)
		currentMax = max(maxEndingHere, currentMax)
	}

	return currentMax
}

func maximumSumWithIndices(arr []int) (int, []int) {
	arrIdxs := []int{0, 0}
	currentMax, maxEndingHere := 0, 0
	for idx := 0; idx < len(arr); idx++ {
		val := arr[idx]
		if maxEndingHere+val < val {
			arrIdxs[0] = idx
			maxEndingHere = val
		} else {
			maxEndingHere = maxEndingHere + val
		}
		if maxEndingHere < currentMax {
			arrIdxs[1] = idx - 1
		} else {
			currentMax = maxEndingHere
			arrIdxs[1] = idx
		}
	}

	return currentMax, arrIdxs
}

func main() {
	fmt.Println(max(-3, 1))
	fmt.Println(maxSum([]int{1, 3, 5, 7, 9}))
	fmt.Println(maxSum([]int{1, -3, 2, 1, -1}))
	fmt.Println(maxSum([]int{1, -3, 2, 1, -1}))
	fmt.Println(maxSum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSum([]int{-1}))
	fmt.Println(maximumSumInefficient([]int{1, -3, 2, 1, -1}))
	fmt.Println(maximumSumInefficient([]int{-1}))
	fmt.Println(maximumSumWithIndices([]int{1, -3, 2, 1, -1}))
	fmt.Println(maximumSumWithIndices([]int{1, -3, 2, 1, -1, -9, 9, -9}))
}
