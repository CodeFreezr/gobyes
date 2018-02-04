/*
 * Write a function which takes an integer and halves it and
 * returns `true` if it was even or `false` if it was odd.
 * For example `half(1)` should return `(0, false)` and 
 * `half(2)` should return `(1, true)`.
 */

package main

import "fmt"

func main() {
	fmt.Println(half(255))
}

func half(num int) (int, bool) {
	halfling := num / 2
	even := false

	if halfling%2 == 0 {
		even = true
	} else {
		even = false
	}
	return halfling, even
}
