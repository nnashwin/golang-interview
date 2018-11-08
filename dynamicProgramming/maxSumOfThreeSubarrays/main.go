package main

import "fmt"

func createDefaultIntSlice(l int) (s []int) {
	for i := 0; i < l; i++ {
		s = append(s, 0)
	}

	return
}

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	sum := 0
	win := createDefaultIntSlice(len(nums) - k + 1)
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Printf("before addition: %d\n", sum)
		sum += nums[i]
		fmt.Printf("after addition: %d\n", sum)
		if i >= k {
			sum -= nums[i-k]
		}
		fmt.Printf("after subtraction: %d\n", sum)
		if i >= k-1 {
			win[i-(k-1)] = sum
		}
	}
	fmt.Println(win)
	fmt.Println(sum)

	j := 0
	left := createDefaultIntSlice(len(win))

	for i := 0; i < len(win); i++ {
		if win[i] > win[j] {
			j = i
		}
		left[i] = j
	}

	j = len(win) - 1
	right := createDefaultIntSlice(len(win))

	for i := len(win) - 1; i >= 0; i-- {
		if win[i] >= win[j] {
			j = i
		}
		right[i] = j
	}

	max := []int{-1, -1, -1}
	for b := k; b < len(win)-k; b++ {
		a := left[b-k]
		c := right[b+k]
		if max[0] == -1 || win[a]+win[b]+win[c] > win[max[0]]+win[max[1]]+win[max[2]] {
			max[0] = a
			max[1] = b
			max[2] = c
		}
	}
	return max
}

func main() {
	fmt.Println(maxSumOfThreeSubarrays([]int{1, 2, 1, 2, 6, 7, 5, 1}, 2))
}
