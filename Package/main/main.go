package main

import (
	"GoLang/Package/stringUtil"
	"fmt"
)

var age = 36

func main() {
	fmt.Println("Imported MyName from stringUtil local package is ", stringUtil.MyName, age)
}
