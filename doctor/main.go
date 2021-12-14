package main

import (
	"bufio"
	"doctor"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput,"\r\n", "", -1)
		userInput = strings.Replace(userInput,"\n","", -1)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}
	}
}