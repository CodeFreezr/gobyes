package main

import "fmt"
import "time"

// https://golang.org/pkg/time/#Tick
//
// func Tick(d Duration) <-chan Time
//
// Tick is a convenience wrapper for NewTicker providing access to the ticking channel only.
// While Tick is useful for clients that have no need to shut down the Ticker,
// be aware that without a way to shut it down the underlying Ticker cannot be recovered by the garbage collector; it "leaks".

func main() {

	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		<-limiter // wait ticks
		fmt.Println("request", req, time.Now())
	}
	fmt.Println("")

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		// use Ticker control time interval
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()


	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
