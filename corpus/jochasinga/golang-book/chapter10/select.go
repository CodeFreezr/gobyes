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
	// bi-directional channel
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
