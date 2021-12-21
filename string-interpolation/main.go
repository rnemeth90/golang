package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            float64
	FavoriteNumber float64
	OwnsADog       bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User
	user.UserName = readString("What is your name?")
	user.Age = readFloat("What is your age?")
	user.FavoriteNumber = readFloat("What is your favorite number?")
	user.OwnsADog = readBool("Do you have a dog? Enter Y or N")

	fmt.Printf("Your name is %s and you are %.f years old. Your favorite number is %.2f. ",
		user.UserName,
		user.Age,
		user.FavoriteNumber,
	)

	if user.OwnsADog {
		fmt.Println("You said you do own a dog.")
	} else {
		fmt.Println("You said you do not own a dog.")
	}
}

func prompt() {
	fmt.Print("> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "" {
			fmt.Println("Please enter a string")
		} else {
			return userInput
		}

	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("Please enter a whole number.")
		} else {
			return num
		}
	}
}

func readBool(s string) bool {
	fmt.Println(s)
	prompt()

	userInput, err := reader.ReadString('\n')

	userInput = strings.Replace(userInput, "\n", "", -1)
	userInput = strings.Replace(userInput, "\r\n", "", -1)

	if err != nil {
		log.Fatal(err)
	}

	if userInput == "Y" {
		return true
	} else {
		return false
	}
}
