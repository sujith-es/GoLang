package main

import (
	"fmt"
	"os"
)

func main() {

	temp, err := os.Open("c:\\test.txt")
	if err != nil {
		fmt.Println("\n Error returned was: ", err)
	} else {
		fmt.Println(temp)
	}
}
