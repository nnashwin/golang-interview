package main

import "fmt"

func longestIncreasingSubsequence(arr []int) int {
	longestSubsequence := []int{}
	longestLen := 0

	for i := 0; i < len(arr); i++ {
		fmt.Println(longestSubsequence)
		num := arr[i]
		if len(longestSubsequence) == 0 || num > longestSubsequence[len(longestSubsequence)-1] {
			longestSubsequence = append(longestSubsequence, num)
		} else {
			idx := 0
			longSubLength := len(longestSubsequence) - 1

			// use binary search to see if any numbers in the already assembled sequence array are less than the current number
			// if they are, keep binary searching to either find a number that is bigger, make the length where that number is and swap it
			// or keep searching until you do not find a number that is bigger (then we will append it to the end)
			for idx < longSubLength {
				fmt.Println("longSubLength: ", longSubLength)
				fmt.Println("num: ", num)
				fmt.Println("idx: ", idx)
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
	fmt.Println(longestSubsequence)
	return longestLen
}

func main() {
	fmt.Println(longestIncreasingSubsequence([]int{2, 3, 4, 7, 13, 17, 9, 16}))
}
