package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = " and don't print the number and press ENTER when ready."

func main() {
	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer int

	fmt.Println("Guess The Number Game....")
	fmt.Println("Think a number between 1 to 10 ", prompt)

	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')

	fmt.Println("Please multiply your number by ", firstNumber, prompt)
	_, _ = reader.ReadString('\n')

	fmt.Println("Please multiply your number by ", secondNumber, prompt)
	_, _ = reader.ReadString('\n')

	fmt.Println("Please divide the result by initial number", prompt)
	_, _ = reader.ReadString('\n')

	fmt.Println("Please subtract the result by ", subtraction, prompt)
	_, _ = reader.ReadString('\n')

	answer = firstNumber*secondNumber - subtraction
	fmt.Println("Your result is: ", answer)
}
