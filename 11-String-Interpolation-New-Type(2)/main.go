package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	HaveADog       bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User
	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = whatIsYourFavoriteNumber()
	user.HaveADog = doYouHaveADog()

	fmt.Println(fmt.Sprintf("Your name is %s, and you are %d years old. Your favorite number is %.2f. You have a dog: %t", user.UserName, user.Age, user.FavoriteNumber, user.HaveADog))
}

func prompt() {
	fmt.Print("-> ")
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

func whatIsYourFavoriteNumber() float64 {
	for {
		fmt.Println("What is your favorite number?")
		prompt()
		age, _ := reader.ReadString('\n')
		age = strings.Replace(age, "\n", "", -1)
		yourAge, err := strconv.ParseFloat(age, 64)
		if err != nil {
			fmt.Println("Enter the correct number...")
		} else {
			return yourAge
		}
	}
}

func doYouHaveADog() bool {
	for {
		fmt.Println("Do you have a dog (y OR n)?")
		prompt()
		haveDog, _ := reader.ReadString('\n')
		haveDog = strings.Replace(haveDog, "\n", "", -1)

		if haveDog == "y" || haveDog == "Y" {
			return true
		} else {
			return false
		}
	}
}
