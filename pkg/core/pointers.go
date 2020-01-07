package main

import (
	"fmt"
	// "runtime"
	// "reflect"
)

// By default go passes by value. Use pointers to pass by reference

func main() {

	name := "Brian"
	ptr := &name 	// pointing to name variable, memory location, now we can pass by reference

	fmt.Println("Name is ", name)
	fmt.Println("Pointer value is ", ptr)
	fmt.Println("Pointer dereferenced is ", *ptr)
	fmt.Println("\n\n")

	// passing by value
	passByValue(name)	
	fmt.Println("Name after pass by value function ", name)

	// passing the pointer to function
	passByReference(&name)
	fmt.Println("Name after pass by reference function ", name)
}

// paramter of type name which is a pointer to a string and returning a string
func passByReference(name *string) string {
	// We have to use the * to dereference the pointer 
	*name = "John"
	fmt.Println("Name changed in passByReference to ", *name)
	fmt.Println("\n\n")
	return *name
}

func passByValue(name string) string {
	name = "Dave"
	fmt.Println("Name changed in passByValue to ", name)
	return name
}
	
	
	
	
	
	
