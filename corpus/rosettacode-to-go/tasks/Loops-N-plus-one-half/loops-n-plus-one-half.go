package main

import "fmt"

func main() {
	for i := 1; ; i++ {
		fmt.Print(i)
		if i == 10 {
			fmt.Println()
			break
		}
		fmt.Print(", ")
	}
}

//\Loops-N-plus-one-half\loops-n-plus-one-half.go
