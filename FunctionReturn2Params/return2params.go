package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "function return 2 parameter"
	author := "sujith surendran"

	fmt.Println(converter(module, author))
}

func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)
	return module, author
}
