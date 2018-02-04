package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plus3(a, b, c int) int {
	return a + b + c
}

// pass array copy to a function
func arrayArgs(a [5]int) {
	a[1] = 88
	fmt.Println("in arrayArgs, a:", a)
}

func sliceArgs(s []int) {
	s[1] = 88
	fmt.Println("in sliceArgs, s:", s)
}

func main() {

	fmt.Println("1+2 =", plus(1, 2))

	fmt.Println("1+2+3 =", plus3(1, 2, 3))

	a := [5]int {1, 2, 3, 4, 5}
	arrayArgs(a)
	fmt.Println("after arrayArgs:", a)

	s := a[:]
	sliceArgs(s)
	fmt.Println("after sliceArgs:", s)
}
