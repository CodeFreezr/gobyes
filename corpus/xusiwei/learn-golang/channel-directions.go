package main

import "fmt"

// When using channels as function parameters,
// you can specify if a channel is meant to only send or receive values.
// This specificity increases the type-safety of the program.

//         send channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//          receive channel        send channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	msg += " passed"
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	go ping(pings, "message")
	go pong(pings, pongs)

	fmt.Println(<-pongs)
}
