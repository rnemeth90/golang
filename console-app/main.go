package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	e := keyboard.Open()

	if e != nil {
		log.Fatal(e)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappacino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Machiatto"
	coffees[6] = "Expresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Machiatto")
	fmt.Println("6 - Expresso")
	fmt.Println("Q - Quit")

	for {
		char, _, e := keyboard.GetSingleKey()

		if e != nil {
			log.Fatal(e)
		}

		i, _ := strconv.Atoi(string(char))

		if char == 'q' || char == 'Q' {
			break
		} else {
			fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
		}

	}
}
