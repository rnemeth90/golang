package main

import "fmt"

func main() {
loop:
	for i := 0; i <= 10; i++ {
		switch i {
		case 1, 2, 3, 4, 5:
			fmt.Println(i, "is less than 6")
		case 6, 7, 8, 9, 10:
			fmt.Println(i, "is above the control. breaking.")
			break loop
		}
	}
}
