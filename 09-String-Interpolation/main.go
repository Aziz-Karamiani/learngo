package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)
	userName := readString("What is your name?")
	age := readInt("How old are you?")

	fmt.Println(fmt.Sprintf("Your name is %s, and you are %d years old.", userName, age))
}

func prompt() {
	fmt.Print("->")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		yourName, _ := reader.ReadString('\n')
		yourName = strings.Replace(yourName, "\n", "", -1)

		if yourName != "" {
			return yourName
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		age, _ := reader.ReadString('\n')
		age = strings.Replace(age, "\n", "", -1)
		yourAge, err := strconv.Atoi(age)
		if err != nil {
			fmt.Println("Enter the correct number...")
		} else {
			return yourAge
		}
	}
}
