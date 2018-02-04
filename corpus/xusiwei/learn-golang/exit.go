package main

import "os"
import "fmt"

func main() {

	// defers will not be run when using os.Exit,
	// so this fmt.Println will never be called.
	defer fmt.Println("!")


	os.Exit(3)
}

// $ go build exit.go
// $ ./build
// $ echo $?
// 3
