package main

import "fmt"

func main() {
	fmt.Printf("My number is %d", Add(4, 5, 6))

}

func Add(a, b, c int) int {
	return a + b + c
}
