package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	dirname := "."

	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
		}

		//log.Println("log")
	}

	//out, err := exec.Command("sloc", "").Output()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//println(out)

	cmd := "sloc"

	if err := exec.Command(cmd).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//	cmnd := exec.Command("cmd", "sloc")
	//	cmnd.Run() // and wait
	//cmnd.Start()
}
