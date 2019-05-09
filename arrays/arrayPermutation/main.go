package main

import "fmt"

func createEmptySlice(len int) (returnArr []string) {
	for i := 0; i < len; i++ {
		returnArr = append(returnArr, "0")
	}
	return
}

func returnPermutationNewArray(arr []string, permArr []int) []string {
	rArr := createEmptySlice(len(arr))
	for i := 0; i < len(permArr); i++ {
		rArr[permArr[i]] = arr[i]
	}

	return rArr
}

func main() {
	fmt.Println(arr)
	fmt.Println(returnPermutationNewArray([]string{"a", "b", "c", "d"}, []int{2, 1, 3, 0}))
}
