package main

import "fmt"

func build2dSquare(length int) [][]int {
	var square [][]int
	for i := 0; i <= length; i++ {
		square = append(square, []int{})
		for j := 0; j <= length; j++ {
			square[i] = append(square[i], 0)
		}
	}

	return square
}

func getHighestGold(mine [][]int) int {
	dp := build2dSquare(len(mine))
	fmt.Println(dp)
	return 0

}

func main() {
	fmt.Println(getHighestGold([][]int{[]int{1, 3, 3}, []int{2, 1, 4}, []int{0, 6, 4}}))
}
