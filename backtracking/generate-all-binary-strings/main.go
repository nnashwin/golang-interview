package main

import "fmt"

func create1dArr(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, 0)
	}

	return arr
}

func printArray(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Printf("\n")
}

func generateAllBinStrings(n int, arr []int, i int) {
	if i == n {
		printArray(arr, n)
		return
	}

	arr[i] = 0
	generateAllBinStrings(n, arr, i+1)

	arr[i] = 1
	generateAllBinStrings(n, arr, i+1)
}

func returnAllBinStrings(n int) {
	arr := create1dArr(n)
	generateAllBinStrings(n, arr, 0)
}

func main() {
	returnAllBinStrings(4)
	returnAllBinStrings(5)
	returnAllBinStrings(6)
}
