package main

import (
	"bufio"
	"fmt"
	"github.com/Aziz-Karamiani/Udemy_Learn_Go_for_Beginners_Crash_Course_Golang/01-HelloWorld/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		}

		fmt.Println(doctor.Response(userInput))
	}
}
