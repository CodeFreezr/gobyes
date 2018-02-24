package main

import (
	"log"

	gc "code.google.com/p/goncurses"
)

func main() {
	_, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer gc.End()
	gc.FlushInput()
}

//\Keyboard-input-Flush-the-keyboard-buffer\keyboard-input-flush-the-keyboard-buffer.go
