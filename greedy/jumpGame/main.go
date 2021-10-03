package main

import "fmt"

func canJump(nums []int) bool {
	n := len(nums)
	last := n - 1
	for i := n - 2; i >= 0; i-- {
		if i + nums[i] >= last {
			last = i
		}
	}

	return last == 0
}

func main() {
	fmt.Println(canJump([]int{2,3,1,1,4}))
	fmt.Println(canJump([]int{3,2,1,0,4}))
}