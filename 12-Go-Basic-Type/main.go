package main

import "log"

// Basic types (Integers, Floats, Strings)
var myInt int = 12
var myInt16 int16 = 13
var myInt32 int32 = 14
var myInt64 int64 = 15
var myUInt uint = 16

var myFloat32 float32 = 12.5
var myFloat64 float64 = 123.5

func main() {
	log.Println(myInt, myInt16, myInt32, myInt64, myUInt)

	log.Println(myFloat32, myFloat64)

	var myString string
	myString = "Aziz"

	log.Println(myString)

	myString = "Karamiani"
}
