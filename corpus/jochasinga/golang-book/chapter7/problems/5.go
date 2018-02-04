package main

import "fmt"

func fib(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(fib(uint(i)))
	}
}
