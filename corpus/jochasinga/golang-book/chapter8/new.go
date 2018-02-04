// Another way to get a pointer is to use the built-in `new` function

package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	// this is a pointer
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}
