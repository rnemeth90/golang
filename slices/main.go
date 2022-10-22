package main

import "fmt"

func main() {
	numSlice := []int{0, 1, 2, 3, 4, 5, 6}
	newSlice := numSlice[:]

	fmt.Println("Value of newSlice[0]:", newSlice[0])
	fmt.Println("Value of numSlice[0]:", numSlice[0])
	fmt.Println("Changing newSlice[0]")
	newSlice[0] = 3
	fmt.Println("Value of newSlice[0]:", newSlice[0])
	fmt.Println("Value of numSlice[0]:", numSlice[0])

}
