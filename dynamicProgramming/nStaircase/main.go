package main

import "fmt"

// first iteration without adding dynamic programming table
// The space complexity of this will be high because it is re-figuring out the same
// sub-problems multiple times
func climbStairs(n int) int {
	// the case where someone jumped down too far
	if n < 0 {
		return 0
	}

	// the case where there is a complete solution, add one to the total number of solutions
	if n == 0 {
		return 1
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairsDynamic(n int, solutionTable map[int]int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	// adds check in map to see if the solution has already been found before
	if _, ok := solutionTable[n]; ok {
		return solutionTable[n]
	}

	solutionTable[n] = climbStairsDynamic(n-1, solutionTable) + climbStairsDynamic(n-2, solutionTable)

	return solutionTable[n]
}

func climbStairsVariableJumps(n int, stepCount []int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	total := 0

	// use a for loop through the steps array to calculate for each step
	// aggregate to a total and return the total on each step
	for i := len(stepCount) - 1; i >= 0; i-- {
		total += climbStairsVariableJumps(n-stepCount[i], stepCount)
	}

	return total
}

func climbStairsVariableJumpsDynamic(n int, stepCount []int, stepTable map[int]int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	if _, ok := stepTable[n]; ok {
		return stepTable[n]
	}

	// prevent assignment issues by initializing the first value at 0 (although I think it will be 0 anyways)
	stepTable[n] = 0
	for i := len(stepCount) - 1; i >= 0; i-- {
		stepTable[n] += climbStairsVariableJumpsDynamic(n-stepCount[i], stepCount, stepTable)
	}

	return stepTable[n]
}

func main() {
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairsDynamic(4, make(map[int]int)))
	fmt.Println(climbStairsVariableJumps(4, []int{1, 2}))
	fmt.Println(climbStairsVariableJumpsDynamic(4, []int{1, 2}, make(map[int]int)))

}
