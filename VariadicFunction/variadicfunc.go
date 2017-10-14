package main

import "fmt"

func main() {
	bestFinish := bestLeagueFinishes(23, 53, 44, 26, 75, 76, 42, 13, 49, 30, 10)
	fmt.Println(bestFinish)
}

func bestLeagueFinishes(finishes ...int) int {

	best := finishes[0]

	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}
