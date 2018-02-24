package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var h, w int
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	d, _ := cmd.Output()
	fmt.Sscan(string(d), &h, &w)
	fmt.Println(h, w)
}

//\Terminal-control-Dimensions\terminal-control-dimensions-2.go
