package main

import (
	"fmt"
	"time"
)

// c channel can only be sent to
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// c channel can both be sent to or receiving from
// thus bi-directional
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// c channel can only be receiving from
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
