package main

import "fmt"

// Reference types (Pointers, Slice, Map, Function, Channel)
// Pointers ??
// use & to get the address
// use * to get the value of Pointers

func main() {
	x := 10
	myPointers := &x

	fmt.Println("x is:", x)
	fmt.Println("myPointers is:", myPointers)

	*myPointers = 15
	fmt.Println("x is now:", x)

	changeVariableUsingPointers(&x)
	fmt.Println("After function call, x is now:", x)
}

func changeVariableUsingPointers(number *int) {
	*number = 25
}
