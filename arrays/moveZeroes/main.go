package main

import "fmt"

// [1, 2, 0, 3, 0, 5]
// [1, 2, 3, 5, 0, 0]

func moveZeroesToEnd(arr []int) []int {
	i := 0
	for j := 0; j < len(arr); j++ {
		if arr[j] != 0 {
			arr[i] = arr[j]
			i++
		}
	}
	for ; i < len(arr); i++ {
		arr[i] = 0
	}

	return arr
}

// [1, 2, 0, 3, 0, 5]
// [0, 0, 1, 2, 3, 5]

func moveZeroesToBeginning(arr []int) []int {
	i := len(arr) - 1
	for j := len(arr) - 1; j >= 0; j-- {
		if arr[j] != 0 {
			arr[i] = arr[j]
			i--
		}
	}
	for ; i >= 0; i-- {
		arr[i] = 0
	}

	return arr
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(moveZeroesToEnd([]int{1, 2, 0, 3, 0, 5}))
	fmt.Println(moveZeroesToBeginning([]int{1, 2, 0, 3, 0, 5}))
}
