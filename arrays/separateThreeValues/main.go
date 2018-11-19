package main

import "fmt"

func sortTripleArrayFirst(arr []string) []string {
	if len(arr) == 0 {
		return arr
	}
	r := 0
	b := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		if arr[i] == "R" {
			arr[i], arr[r] = arr[r], arr[i]
			r++
		}

	}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(b)
		fmt.Println(i)
		if arr[i] == "B" {
			arr[i], arr[b] = arr[b], arr[i]
			b--
		}
	}

	return arr
}

func sortTripleArraySecond(arr []string) []string {
	if len(arr) == 0 {
		return arr
	}
	i, j := 0, 0
	n := len(arr) - 1

	for j <= n {
		if arr[j] == "R" {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		} else if arr[j] == "B" {
			arr[n], arr[j] = arr[j], arr[n]
			n--
		} else {
			j++
		}
	}

	return arr
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(sortTripleArrayFirst([]string{"G", "B", "R", "R", "B", "R", "G"}))
	fmt.Println(sortTripleArraySecond([]string{"G", "B", "R", "R", "B", "R", "G"}))
}
