package main

import (
	"fmt"
	"runtime"
)

// function main. Prints hello world text, OS and Go Arch.
func main() {
	fmt.Println("Hello world.", runtime.GOOS, runtime.GOARCH)
}
