package main

import "fmt"

func withinBounds(point []int, screen [][]string) bool {
	row, col := point[0], point[1]

	rowHeight := len(screen)
	colWidth := len(screen[0])

	return row >= 0 && row < rowHeight && col >= 0 && col < colWidth
}

func paintFill(point []int, screen [][]string, desiredCol string) {
	row, col := point[0], point[1]

	if withinBounds(point, screen) && screen[row][col] != desiredCol {
		fmt.Println(screen[row][col])
		screen[row][col] = desiredCol
		fmt.Println(screen[row][col])

		if withinBounds([]int{row + 1, col}, screen) && screen[row+1][col] != desiredCol {
			paintFill([]int{row + 1, col}, screen, desiredCol)
		}

		if withinBounds([]int{row, col + 1}, screen) && screen[row][col+1] != desiredCol {
			fmt.Println("next col same row")
			paintFill([]int{row, col + 1}, screen, desiredCol)
		}

		if withinBounds([]int{row + 1, col + 1}, screen) && screen[row+1][col+1] != desiredCol {
			paintFill([]int{row + 1, col + 1}, screen, desiredCol)
		}

		if withinBounds([]int{row - 1, col - 1}, screen) && screen[row-1][col-1] != desiredCol {
			paintFill([]int{row - 1, col - 1}, screen, desiredCol)
		}

		if withinBounds([]int{row - 1, col}, screen) && screen[row-1][col] != desiredCol {
			paintFill([]int{row - 1, col}, screen, desiredCol)
		}

		if withinBounds([]int{row, col - 1}, screen) && screen[row][col-1] != desiredCol {
			paintFill([]int{row, col - 1}, screen, desiredCol)
		}
	}
}

func main() {
	fmt.Println("Paintfill!!!!")
	screen := [][]string{[]string{"y", "g", "y"}, []string{"y", "g", "y"}, []string{"g", "y", "g"}}
	screen2 := [][]string{[]string{"g", "g", "g"}, []string{"b", "b", "g"}, []string{"b", "b", "g"}}
	screen3 := [][]string{[]string{"r", "g", "g"}, []string{"b", "b", "b"}, []string{"b", "b", "g"}}
	paintFill([]int{0, 1}, screen, "y")
	paintFill([]int{1, 1}, screen2, "r")
	paintFill([]int{0, 1}, screen3, "b")

	fmt.Println(screen)
	fmt.Println(screen2)
	fmt.Println(screen3)
}
