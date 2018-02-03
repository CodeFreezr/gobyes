package main

import "fmt"

func fToM(feet float64) (meter float64) {
	meter = feet * 0.3048
	return
}

func main() {
	fmt.Println(fToM(57))
}
