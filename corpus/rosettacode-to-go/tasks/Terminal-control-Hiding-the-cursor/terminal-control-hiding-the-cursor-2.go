package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("\033[?25l")
	time.Sleep(3 * time.Second)
	fmt.Print("\033[?25h")
	time.Sleep(3 * time.Second)
}

//\Terminal-control-Hiding-the-cursor\terminal-control-hiding-the-cursor-2.go
