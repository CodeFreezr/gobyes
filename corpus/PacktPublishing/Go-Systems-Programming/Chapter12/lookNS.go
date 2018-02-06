package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a domain!")
		os.Exit(100)
	}

	domain := arguments[1]

	NSs, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	for _, NS := range NSs {
		fmt.Println(NS.Host)
	}
}
