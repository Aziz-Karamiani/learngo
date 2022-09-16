package packageone

import "fmt"

var privateNumber = 12
var PublicNumber = 132

func PrintNumber() {
	fmt.Println(PublicNumber)
}

func printNumber() {
	fmt.Println(privateNumber)
}
