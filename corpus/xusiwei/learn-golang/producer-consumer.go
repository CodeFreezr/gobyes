package main

import "fmt"

func producer(queue chan string, num int) {
	for i := 1; i <= num; i++ {
		queue <- fmt.Sprintf("item-%d", i)
	}
	close(queue)
}

func consumer(queue chan string, done chan bool) {
	for item := range queue {
		fmt.Println(item)
	}
	done <- true
}

func main() {
	queue := make(chan string, 3)
	done := make(chan bool)

	go producer(queue, 10)
	go consumer(queue, done)

	<-done
}
