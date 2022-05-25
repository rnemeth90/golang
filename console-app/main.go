package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("->")

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(userInput)
		}
	}
}