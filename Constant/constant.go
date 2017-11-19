package main

import (
	"fmt"
)

const (
	_ = iota
	B = iota * 10
	C = iota * 10
)

const (
	_ = iota
	D = iota * 10
	E = iota * 10
)

func main() {
	const name = "Sujith"
	const course = "Docker deep dive"

	fmt.Println("\nHi", name, " You're currently watching", "'", course, "'", " which is constant.")
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
}
