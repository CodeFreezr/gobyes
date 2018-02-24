package main

import (
	"fmt"
	"math"
)

func main() {
	for i := int8(0); ; i++ {
		fmt.Printf("%o\n", i)
		if i == math.MaxInt8 {
			break
		}
	}
}

//\Count-in-octal\count-in-octal-1.go
