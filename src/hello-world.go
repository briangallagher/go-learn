package main

import (
	"fmt"
	"runtime"
	"reflect"
)

var (
	name, course string // Name of sub
)


func main() {
	module := 10.12314
	ptr := &module		// pointer variable

	fmt.Println("Hello from", runtime.GOOS)
	fmt.Println("name", name)
	fmt.Println("course", reflect.TypeOf(course))
	fmt.Println("module", module)

	fmt.Println("Memory address is ", ptr)
	fmt.Println("Memory Value is ", *ptr)		// dereference a pointer variable
}




