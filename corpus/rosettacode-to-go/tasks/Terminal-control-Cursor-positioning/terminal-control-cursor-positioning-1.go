package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("tput", "-S")
	cmd.Stdin = bytes.NewBufferString("clear\ncup 5 2")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("Hello")
}

//\Terminal-control-Cursor-positioning\terminal-control-cursor-positioning-1.go
