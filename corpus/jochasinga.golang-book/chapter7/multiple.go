package main

import "fmt"

func f() (int, int) {
	return 5, 6
}

func main() {
	x, y := f()
	fmt.Println(x, y)
}
