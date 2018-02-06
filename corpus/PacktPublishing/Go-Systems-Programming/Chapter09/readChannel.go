package main

import (
	"fmt"
	"time"
)

func writeChannel(c chan<- int, x int) {
	fmt.Println(x)
	c <- x
	close(c)
	fmt.Println(x)
}

func main() {
	c := make(chan int)
	go writeChannel(c, 10)
	time.Sleep(2 * time.Second)
	fmt.Println("Read:", <-c)
	time.Sleep(2 * time.Second)

	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}
