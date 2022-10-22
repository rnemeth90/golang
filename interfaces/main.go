package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("testFile")
	b, _ := io.ReadAll(f)
	fmt.Println(string(b))
}
