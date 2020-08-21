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

func checkIdx(dp []int, i int) int {
	fmt.Println("dp: ", dp)
	fmt.Println(" i: ", i)
	if i < 0 {
		return 0
	}

	return dp[i]
}

func robHousesDp(nums []int) int {
	dp := make([]int, len(nums), len(nums))

	m := 0

	for i := 0; i < len(nums); i++ {
		fmt.Println(" nums: ", nums)
		dp[i] = nums[i] + max(checkIdx(dp, i-2), checkIdx(dp, i-3))
		if dp[i] > m {
			m = dp[i]
		}
	}

	return m
}

func main() {
	fmt.Println(robHouse([]int{1, 2, 3, 1}))
	fmt.Println(robHouse([]int{2, 7, 9, 3, 1}))
	fmt.Println(robHousesDp([]int{1, 2, 3, 1}))
	fmt.Println(robHousesDp([]int{2, 7, 9, 3, 1}))
	fmt.Println(robHousesDp([]int{6, 3, 10, 8, 2, 10, 3, 5, 10, 5, 3}))

}
