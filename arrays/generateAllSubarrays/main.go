package main

import "fmt"

func bruteForcePrintAllSubarrays(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := i; k < j; k++ {
				fmt.Printf("%d", arr[k])
			}
			fmt.Println()
		}
	}
}

func verboseOptimizedPrintAllSubarrays(arr []int) {
	var result [][]int
	for i := 0; i < len(arr); i++ {
		var subArrContainer []int
		for j := i; j < len(arr); j++ {
			subArrContainer = append(subArrContainer, arr[j])
			result = append(result, subArrContainer)
		}
	}
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

func bestPrintAllSubarrays(arr []int) {
	for i := 0; i < len(arr); i++ {
		var subArrContainer []int
		for j := i; j < len(arr); j++ {
			subArrContainer = append(subArrContainer, arr[j])
			fmt.Println(subArrContainer)
		}
	}
}

func main() {
	// bruteForcePrintAllSubarrays([]int{1, 3, 5, 7, 9})
	// verboseOptimizedPrintAllSubarrays([]int{1, 3, 5, 7, 9})
	bestPrintAllSubarrays([]int{1, 3, 5, 7, 9})
}
