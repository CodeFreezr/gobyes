package main

import "fmt"


func main() {
	fmt.Println("go" + "lang")

	// integer and float operations
	fmt.Println("5+2 = ", 5+2)
	fmt.Println("5-2 = ", 5-2)
	fmt.Println("5*2 = ", 5*2)
	fmt.Println("5/2 = ", 5/2)
	fmt.Println("5%2 = ", 5%2)
	fmt.Println("5<<2 = ", 5<<2)
	fmt.Println("5>>2 = ", 5>>2)
	fmt.Println("10.0/3.0 = ", 10.0/3.0)

	// Boolean, with the operators as you'd expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
