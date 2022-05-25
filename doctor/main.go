package main

import (
	"fmt"

	"myapp/doctor.go"
)

func main() {
	var whatToSay string
	whatToSay = doctor.Intro()
	fmt.Println(whatToSay)
}
