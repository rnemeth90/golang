package main

import (
	"fmt"
)

func main() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)

	var myPointer *int
	myPointer = &x
	fmt.Println(*myPointer + 15)
}
