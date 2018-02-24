package main

import (
	"log"

	gc "code.google.com/p/goncurses"
)

func main() {
	s, err := gc.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer gc.End()
	s.Move(5, 2)
	s.Println("Hello")
	s.GetChar()
}

//\Terminal-control-Cursor-positioning\terminal-control-cursor-positioning-3.go
