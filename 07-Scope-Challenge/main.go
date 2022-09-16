package main

import "github.com/Aziz-Karamiani/Udemy_Learn_Go_for_Beginners_Crash_Course_Golang/07-Scope-Challenge/packageone"

var myVar = "My Package Level Variable"

func main() {
	var blockVar = "Block Variable in main function"
	packageone.PrintMe(myVar, blockVar)
}
