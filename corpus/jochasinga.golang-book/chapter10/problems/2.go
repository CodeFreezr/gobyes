// Write your own `Sleep` function using `time.After`

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

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case <-time.After(time.Second * 2):
				fmt.Println("timeout")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
