package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello world.", runtime.GOOS, runtime.GOARCH)
}
