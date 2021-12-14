package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "press ENTER when ready."

// guess the number game

func main() {
	// seed the random num generator
	rand.Seed(time.Now().Unix())

	var firstNumber int = rand.Intn(8) + 2
	var secondNumber int = rand.Intn(8) + 2
	var subtraction int = rand.Intn(8)+ 2
	answer := firstNumber * secondNumber - subtraction

	gameRunTime(firstNumber,secondNumber,subtraction,answer)

}

func gameRunTime(firstNumber int, secondNumber int, subtraction int, answer int){
	reader := bufio.NewReader(os.Stdin)

	// display welcome/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")
	fmt.Println("Think of a number between 1 and 10.")
	fmt.Println(prompt)
	reader.ReadString('\n')

	// go through the game
	fmt.Println("Multiply your number by", firstNumber, "and", prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply by the result by", secondNumber, "and", prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of and ", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, "and", prompt)
	reader.ReadString('\n')

	// display the answer
	fmt.Println("Your number is:",answer)
}
