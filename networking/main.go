package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	dnsName, e := net.LookupIP(name)

	if e != nil {
		log.Fatal(e)
	}

	// addr := net.ParseIP(name)
	if dnsName == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", dnsName)
	}
	os.Exit(0)
}
