package main

import (
	"bufio"
	"fmt"
	"os"
)

// var mystring string = "Hello Ryan"
var myint int = 20

// this is a comment

/*
a multiline comment
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	result, _ := reader.ReadString('\n')
	sayHelloWorld(result,myint)
}

func sayHelloWorld(whatToSay string,howManyTimes int){
	for i := 0; i < howManyTimes; i++ {
		fmt.Println(whatToSay)

	}
}


