package main

import (
	"fmt"
)

func permuteUnique(nums []int) [][]int {
	permutations := [][]int{}
	backtrack(nums, &nums, &permutations)
	return permutations
}

func backtrack(nums []int, raw *[]int, permutations *[][]int) {
	if len(nums) <= 1 {
		fmt.Println("raw: ", *raw)
		*permutations = append(*permutations, duplicate(*raw))
		return
	}

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		fmt.Println("nums: ", nums)
		fmt.Println("nums[0]: ", nums[0])
		fmt.Println("nums[i]: ", nums[i])

		fmt.Println("raw before swap: ", raw)
		nums[0], nums[i] = nums[i], nums[0]
		fmt.Println("raw after swap: ", raw)
		backtrack(nums[1:], raw, permutations)
		nums[0], nums[i] = nums[i], nums[0]
		m[nums[i]] = i
	}
}

func duplicate(a []int) []int {
	ret := make([]int, len(a))
	copy(ret, a)
	return ret
}

func main() {
	fmt.Println(permuteUnique([]int{1, 2, 2, 1}))
}
