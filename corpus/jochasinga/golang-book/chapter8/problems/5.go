package main

import "fmt"

func swap(x *int, y *int) {
	xVal := *x
	yVal := *y

	*x = yVal
	*y = xVal
}

func main() {
	x := 8
	y := 3
	swap(&x, &y)
	fmt.Println("x Should be 3 ->", x)
	fmt.Println("y Should be 8 ->", y)
}
