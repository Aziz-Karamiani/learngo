package main

import (
	"fmt"
)

// Aggregate Types (array && struct)

type Car struct {
	NumberOfDoors int
	NumberOfTiers int
	Model         string
	Year          int
	Factory       string
}

func main() {
	var myString [3]string

	myString[0] = "Dog"
	myString[1] = "Cat"
	myString[2] = "Donkey"

	fmt.Printf("Your favorite animals are: %s %s %s", myString[0], myString[1], myString[2])
	fmt.Println()

	car := Car{
		NumberOfDoors: 4,
		NumberOfTiers: 4,
		Model:         "Ford",
		Year:          2020,
		Factory:       "Ford",
	}

	fmt.Printf("You have a %d %s car.", car.Year, car.Model)

}
