package main

import (
	"fmt"
	"github.com/Aziz-Karamiani/Udemy_Learn_Go_for_Beginners_Crash_Course_Golang/06-Scope/packageone"
)

var one = "One"

func main() {
	fmt.Println(one)
	var one = "This is another variable"
	fmt.Println(one)
	printOne()

	fmt.Println(packageone.PublicNumber)
	packageone.PrintNumber()
}

func printOne() {
	var one = "This is another variable 1"
	fmt.Println(one)
}
