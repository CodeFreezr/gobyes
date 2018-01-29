package main

import "fmt"

// likes Python's tuple
func foo() (int, int) {
	return 3, 7
}

func main() {

	a, b := foo()
	fmt.Println("a, b:", a, b)

	_, c := foo()
	fmt.Println("c:", c)

	d, _ := foo()
	fmt.Println("d:", d)
}
