package main

import (
	"fmt"
)

func main() {

	type courseMeta struct {
		Name   string
		Level  string
		Rating float64
	}

	dockerDeepDive := courseMeta{
		Name:   "Docker Deep Dive",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println("\nCourse Name:", dockerDeepDive.Name)
	fmt.Println("\nCourse Level:", dockerDeepDive.Level)
	fmt.Println("\nCourse Rating:", dockerDeepDive.Rating)

}
