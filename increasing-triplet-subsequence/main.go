package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	min := nums[0]
	second_min := int(math.MaxInt32)

	for i := 1; i < len(nums); i++ {
		// This will set the absolute minimum number in the array to min
		if nums[i] <= min {
			min = nums[i]
			// If the minimum number is set, this will set the second-minimum number to the second minimum
		} else if nums[i] <= second_min {
			second_min = nums[i]
			// If both of the top ifs do not occur, which means that the first and second minimums have already been set and there is a third number that is bigger than them,
			// return true.  We have found an increasing triplet
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 6, 3}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 6, 7}))
	fmt.Println(increasingTriplet([]int{5, 2, 3, 1, 7}))
	fmt.Println(increasingTriplet([]int{1, 2, 3, 5, 6, 7}))
}
