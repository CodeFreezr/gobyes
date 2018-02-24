package main

import "fmt"

func printOnes() {
	for {
		fmt.Println("1")
	}
}

func main() {
	go printOnes()
	for {
		fmt.Println("0")
	}
}

//\Flow-control-structures\flow-control-structures-3.go
