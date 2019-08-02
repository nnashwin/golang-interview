package main

import "fmt"

func buildMatrix(n int) (maze [][]int) {
	for i := 0; i < n; i++ {
		maze = append(maze, []int{})
		for j := 0; j < n; j++ {
			maze[i] = append(maze[i], 0)
		}
	}

	return
}

func isSafe(maze [][]int, x int, y int) bool {
	return x >= 0 && y >= 0 && x < len(maze) && y < len(maze) && maze[x][y] != 0
}

func findPath(maze [][]int, x int, y int, solution *[][]int) bool {
	l := len(maze)

	if x == l-1 && y == l-1 {
		(*solution)[x][y] = 1
		return true
	}

	if isSafe(maze, x, y) == true {
		(*solution)[x][y] = 1
		if findPath(maze, x+1, y, solution) {
			return true
		}

		if findPath(maze, x, y+1, solution) {
			return true
		}

		(*solution)[x][y] = 0
		return false
	}

	return false
}

func findPathWithBackMotion(maze *[][]int, x int, y int, solution *[][]int) bool {
	l := len(*maze)

	if x == l-1 && y == l-1 {
		(*solution)[x][y] = 1
		return true
	}

	(*maze)[x][y] = 0
	(*solution)[x][y] = 1

	if isSafe((*maze), x+1, y) && findPathWithBackMotion(maze, x+1, y, solution) {
		return true
	}

	if isSafe(*maze, x, y+1) && findPathWithBackMotion(maze, x, y+1, solution) {
		return true
	}

	if isSafe(*maze, x-1, y) && findPathWithBackMotion(maze, x-1, y, solution) {
		return true
	}

	if isSafe(*maze, x, y-1) && findPathWithBackMotion(maze, x, y-1, solution) {
		return true
	}

	(*solution)[x][y] = 0
	return false
}

func ratInAMaze(maze [][]int) [][]int {
	solution := buildMatrix(len(maze))

	if findPath(maze, 0, 0, &solution) == true {
		return solution
	}

	return nil
}

func ratInAMazeBack(maze [][]int) [][]int {
	solution := buildMatrix(len(maze))
	if findPathWithBackMotion(&maze, 0, 0, &solution) == true {
		return solution
	}

	return nil
}

func main() {
	maze := [][]int{[]int{1, 0, 0, 0},
		[]int{1, 1, 1, 1},
		[]int{0, 0, 1, 0},
		[]int{0, 1, 1, 1}}
	fmt.Println(maze)

	fmt.Println(ratInAMaze(maze))

	maze2 := [][]int{
		[]int{1, 1, 0, 0},
		[]int{0, 1, 0, 0},
		[]int{1, 1, 0, 0},
		[]int{1, 1, 1, 1},
	}

	fmt.Println(ratInAMazeBack(maze2))
}
