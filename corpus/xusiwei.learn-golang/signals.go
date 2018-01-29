package main

import "fmt"
import "os"
import "os/signal"
import "syscall"


func main() {

	// Go signal notification works by sending os.Signal values on a channel.
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)


	// signal.Notify registers the given channel to receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)


	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
