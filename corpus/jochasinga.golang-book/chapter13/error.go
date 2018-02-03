package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("error message")
	fmt.Println(err)
}
