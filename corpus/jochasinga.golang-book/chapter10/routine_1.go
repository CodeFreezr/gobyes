package main

import "fmt"

func f(n int) {
	for i := 0; i < 100; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
