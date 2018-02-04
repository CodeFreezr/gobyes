// What is the value of x after running this program?

package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
	fmt.Println(*x)
}
func main() {
	x := 1.5
	square(&x)
}

// The answer is 2.25
