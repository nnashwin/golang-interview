package main

import "fmt"

func multiplyArraySubtractIndex(input []int) []int {
	total := 1

	for i := 0; i < len(input); i++ {
		total *= input[i]
	}

	for i := len(input) - 1; i >= 0; i-- {
		input[i] = total / input[i]
	}

	return input
}

func main() {
	fmt.Println(multiplyArraySubtractIndex([]int{1, 2, 3, 4, 5}))
	fmt.Println(multiplyArraySubtractIndex([]int{3, 2, 1}))
}
