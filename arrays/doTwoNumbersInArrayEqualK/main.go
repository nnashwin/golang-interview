package main

import "fmt"

func main() {
	test1 := []int{10, 15, 3, 7} // 17
	test2 := []int{1, 3, 5, 8}   // 22
	test3 := []int{2, 4, 6, 8}   // 14

	result1 := doTwoNumbersInArrayEqualK(test1, 17)
	result2 := doTwoNumbersInArrayEqualK(test2, 22)
	result3 := doTwoNumbersInArrayEqualK(test3, 14)

	fmt.Println("test 1 == 17", result1)
	fmt.Println("test 2 == 22", result2)
	fmt.Println("test 3 == 14", result3)
}

func doTwoNumbersInArrayEqualK(arr []int, k int) bool {
	for i, j := 0, len(arr)-1; i < j; {
		sum := arr[i] + arr[j]
		if sum < k {
			i++
		} else if sum > k {
			j--
		} else {
			return true
		}
	}
	return false
}
