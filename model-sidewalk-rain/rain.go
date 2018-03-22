package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("It's RAINING!!!!")
	sidewalk := [10][10]int{}
	count, iterations := 0, 0

	for count < 100 {
		x := randomInt(0, 10)
		y := randomInt(0, 10)
		point := sidewalk[randomInt(0, 10)][randomInt(0, 10)]
		iterations += 1

		if point == 0 {
			sidewalk[x][y] = 1
			count += 1
		}
	}

	fmt.Println(iterations)

}
