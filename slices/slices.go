package main

import (
	"fmt"
)

func main() {

	fmt.Println("Program to display Array and Slices\n")
	//mySlices := make([]string, 10, 20)
	mySlices := []int{2, 3, 4, 5, 6, 7}
	fmt.Println("Actual mySlices Array")
	fmt.Println(mySlices)
	fmt.Printf("Length is: %d \nCapacity  is: %d", len(mySlices), cap(mySlices))

	mySlices[0] = 1000

	fmt.Println("\nmySlices[0] Array")
	fmt.Println(mySlices)
	fmt.Println("\nmySlices[2:5] Array")
	sliceOfSlice := mySlices[:5]
	fmt.Println(sliceOfSlice)
}
