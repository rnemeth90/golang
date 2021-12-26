package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "google.org:80")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connecting...")
	fmt.Fprintf(conn, "GET / HTTP/1.0")
}
