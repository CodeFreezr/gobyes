package main

import "fmt"
import "time"

// example of using a blocking receive to wait for a goroutine to finish.

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool)

	go worker(done)

	<-done
}
