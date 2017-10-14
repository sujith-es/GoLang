package main

import (
	"fmt"
)

func main() {
	name := "Sujith"
	course := "Docker deep dive"

	fmt.Println("\nHi", name, " You're currently watching", course)

	changeCourse(course)
	fmt.Println("\nModified course value returned by function is", course)
}

func changeCourse(course string) string {

	course = "Google Go language Fundamentals"
	fmt.Println("\nNew course value is ", course)
	return course
}
