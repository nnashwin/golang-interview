package main

import "fmt"

func subsets(nums []int) [][]int {
	powerset := [][]int{}
	backtrack(nums, 0, []int{}, &powerset)
	return powerset
}

func backtrack(nums []int, start int, subset []int, powerset *[][]int) {
	newSubset := make([]int, len(subset)+1)
	copy(newSubset[0:len(subset)], subset)
	*powerset = append(*powerset, newSubset[:len(newSubset)-1])

	for i := start; i < len(nums); i++ {
		subset = append(subset, nums[i])
		backtrack(nums, i+1, subset, powerset)

		subset = subset[:len(subset)-1]
	}
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
