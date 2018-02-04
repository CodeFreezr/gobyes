// When a function is called taking an argument, that argument's value
// is copied to the function:

// func zero(x int) {
// 	x = 0
// }
// func main() {
// 	x := 5
//	zero(x)
//	fmt.Println(x) // x is still 5

// If we want to actually pass in the argument and modify it, not just
// its value, we need to pass in its pointer.

package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0 
}
