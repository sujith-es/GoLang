package main

import (
	"fmt"
)

func main() {

	fmt.Println("Program to display Array and Slices\n")
	mySlices := make([]int, 1, 4)
	//mySlices := []int{2, 3, 4, 5, 6, 7}
	fmt.Println("Actual mySlices Array")
	fmt.Println(mySlices)
	fmt.Printf("Length is: %d \nCapacity  is: %d", len(mySlices), cap(mySlices))

	mySlices[0] = 1000

	fmt.Println("\nmySlices[0] Array")
	fmt.Println(mySlices)
	fmt.Println("\nmySlices[2:5] Array")
	sliceOfSlice := mySlices[:2]
	fmt.Println(sliceOfSlice)

	fmt.Println("\nIncrease capacity of underlying Array used by slice")
	for index := 1; index < 17; index++ {
		mySlices = append(mySlices, index)
		fmt.Printf("\nCapacity is: %d", cap(mySlices))
	}
	newSlices := []int{200, 300, 500}
	mySlices = append(mySlices, newSlices...)
	fmt.Println("\nnewSlices appended")
	fmt.Println(mySlices)
}
