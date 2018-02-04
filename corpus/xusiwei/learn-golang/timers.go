package main

import "fmt"
import "time"

func main() {

	// timer provides a channel that will be notified at that time.
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	// you can cancel the timer before it expires.
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stoped")
	}
}
