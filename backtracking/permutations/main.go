package main

import "fmt"

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func permute(nums []int) [][]int {
	allPermutations := [][]int{}
	backtrack(nums, &allPermutations, []int{})

	return allPermutations
}

func backtrack(nums []int, permutations *[][]int, subset []int) {
	newSubset := make([]int, len(subset))
	copy(newSubset[0:len(subset)], subset)
	if len(newSubset) == len(nums) {
		*permutations = append(*permutations, newSubset)
	} else {
		for i := 0; i < len(nums); i++ {
			if contains(subset, nums[i]) {
				continue
			}
			subset = append(subset, nums[i])
			backtrack(nums, permutations, subset)
			subset = subset[:len(subset)-1]
		}
	}
}

func main() {
	fmt.Println(findPermutations([]int{1, 2, 3}))
}
