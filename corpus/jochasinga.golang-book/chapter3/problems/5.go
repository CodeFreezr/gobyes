// Whats the value of the expression `(true && false) ||
// (false && true) || !(false && false)`?

// true && false = false
// false && true = false
// false && false = false
// Therefore, false || false || !false = true

package main

import "fmt"

func main() {
	fmt.Println((true && false) || (false && true) || !(false && false))
}
