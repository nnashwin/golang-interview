package main

import "fmt"

func createEmptyStrSlice(len int) (returnArr []string) {
	for i := 0; i < len; i++ {
		returnArr = append(returnArr, "")
	}
	return
}

func returnPermutationNewArray(arr []string, permArr []int) []string {
	rArr := createEmptyStrSlice(len(arr))
	for i := 0; i < len(permArr); i++ {
		rArr[permArr[i]] = arr[i]
	}

	return rArr
}

func returnPermutationInPlace(arr []int, permArr []int) []int {
	for i := 0; i < len(permArr); i++ {
		permArr[i] = arr[permArr[i]]
	}

	return permArr
}

func main() {
	fmt.Println(returnPermutationNewArray([]string{"a", "b", "c", "d"}, []int{2, 1, 3, 0}))
	fmt.Println(returnPermutationInPlace([]int{3, 4, 5, 6}, []int{2, 1, 3, 0}))
}
