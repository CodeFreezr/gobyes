package main

import (
	"os"
	"os/exec"
)

func main() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

//\Terminal-control-Clear-the-screen\terminal-control-clear-the-screen-1.go
