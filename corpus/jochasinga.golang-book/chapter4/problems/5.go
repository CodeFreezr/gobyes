package main

import "fmt"

func main() {
	fmt.Println(fToC(57))
}

func fToC(fTemp float32) (cTemp float32) {
	cTemp = (fTemp - 32) * 5 / 9
	return
}
