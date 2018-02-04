package main

import "fmt"

// use for-range syntax to iterate over values received from a channel.

func main() {

	queue := make(chan string, 2)

	go func() {
		queue <- "one"
		queue <- "two"
		close(queue)
	}()

// Because we closed the channel above, 
// the iteration terminates after receiving the 2 elements.
// If we didn’t close it we’d block on a 3rd receive in the loop.
	for elem := range queue {
		fmt.Println(elem)
	}
}
