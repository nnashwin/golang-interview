package main

import "fmt"

func longestIncreasingSubsequence(arr []int) int {
	longestSubsequence := []int{}
	longestLen := 0

	for i := 0; i < len(arr); i++ {
		num := arr[i]
		if len(longestSubsequence) == 0 || num > longestSubsequence[len(longestSubsequence)-1] {
			longestSubsequence = append(longestSubsequence, num)
		} else {
			idx := 0
			longSubLength := len(longestSubsequence) - 1

			for idx < longSubLength {
				mid := (idx + longSubLength) / 2
				if longestSubsequence[mid] < num {
					idx += mid + 1
				} else {
					longSubLength = mid
				}
			}
			longestSubsequence[longSubLength] = num
		}

		longestLen = len(longestSubsequence)
	}
	return longestLen
}

func main() {
	fmt.Println(longestIncreasingSubsequence([]int{9, 1, 3, 7, 5, 6, 20}))
	fmt.Println(longestIncreasingSubsequence([]int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}))
}
