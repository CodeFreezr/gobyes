package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "git"
	cmdArgs := []string{"--version"}

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
		os.Exit(1)
	}
	/*
		sha := string(cmdOut)
		firstSix := sha[:6]
		fmt.Println("The first six chars of the SHA at HEAD in this repo are", firstSix)
	*/
	fmt.Println(cmdOut)
}
