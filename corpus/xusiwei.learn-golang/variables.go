package main

import "fmt"

func main() {
	// declare one variable
	var a string = "initial"
	fmt.Println(a)

	// declare more variables
	var b, c int = 1, 2
	fmt.Println(b, c)

	// infer type by value
	var d = true
	fmt.Println(d)

	// without given initialization are *zero-valued*
	var e int
	fmt.Println(e)

	// `:=` is shorthand for declaring and initializing a variable
	f := "short"
	fmt.Println(f)

// errors:
//	g = "bad"  // undefined
//	f = 1      // assignment type miss match 
}
