package main

import (
	"fmt"
)

func main() {
	coursesInProgress := []string{"angular 4", "Google Go programming", "Blockchain"}

	for _, i := range coursesInProgress {
		fmt.Println(i)
	}
}
