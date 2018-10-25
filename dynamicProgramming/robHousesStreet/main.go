package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robHouse(nums []int) int {
	val := 0
	for i := len(nums); i >= 0; i-- {
		temp := robHouses(nums[:i], make(map[int]int))
		val = max(temp, val)
	}

	return val
}

func robHouses(nums []int, alreadyRobbed map[int]int) int {
	lastIndex := len(nums) - 1
	// check for 0
	if len(nums) < 1 || nums[lastIndex] == 0 {
		return 0
	}

	if _, ok := alreadyRobbed[lastIndex]; ok {
		return alreadyRobbed[lastIndex]
	}

	minus2 := len(nums) - 2
	if minus2 < 0 {
		minus2 = 0
	}

	minus3 := len(nums) - 3
	if minus3 < 0 {
		minus3 = 0
	}

	alreadyRobbed[lastIndex] = nums[lastIndex] + max(robHouses(nums[:minus2], alreadyRobbed), robHouses(nums[:minus3], alreadyRobbed))

	return alreadyRobbed[lastIndex]

}

func main() {
	fmt.Println(robHouse([]int{1, 2, 3, 1}))
	fmt.Println(robHouse([]int{2, 7, 3, 9, 1}))
	fmt.Println(robHouse([]int{6, 3, 10, 8, 2, 10, 3, 5, 10, 5, 3}))

}
