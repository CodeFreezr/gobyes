package main

import (
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func main() {
	h := ripemd160.New()
	h.Write([]byte("Rosetta Code"))
	fmt.Printf("%x\n", h.Sum(nil))
}

//\RIPEMD-160\ripemd-160.go
