package main

import (
	"fmt"
	"reflect"
)

var (
	name, course, module = "Sujith", "Docker deep dive", 3.2
)

func main() {
	prt := &module
	fmt.Println("Name is ", name, " and is of type ", reflect.TypeOf(name))
	fmt.Println("module is ", module, " and is of type ", reflect.TypeOf(module))
	fmt.Println("memory address of *module* is ", prt, " and value is ", *prt)
}
