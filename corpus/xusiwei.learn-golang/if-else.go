package main

import "fmt"

func main() {
	// if
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// if-else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if-elseif-else
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digits")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
