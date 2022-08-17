package main

import (
	"fmt"
	"github.com/Aziz-Karamiani/Udemy_Learn_Go_for_Beginners_Crash_Course_Golang/01-HelloWorld/doctor"
)

func main() {
	var whatToSay string
	whatToSay = doctor.Intro()
	fmt.Println(whatToSay)
}
