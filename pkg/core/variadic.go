package main

import (
	"fmt"
)

func main() {
	bestFinish := bestLeagueFinish(12, 10, 11, 9, 3, 2)
	fmt.Println(bestFinish)
}

// Pass in a variable number of parameters
func bestLeagueFinish(finishes ...int) int {

	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}

