package main

import "fmt"

// Closing a channel indicates that no more values will be sent on it.
// This can be useful to communicate completion to the channel’s receivers.

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// the more value will be false if jobs has been closed
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close (jobs)
	fmt.Println("sent all jobs")

	<-done
}
