package main

import (
	"fmt"
	"golang.org/x/crypto/md4"
)

func main() {
	h := md4.New()
	h.Write([]byte("Rosetta Code"))
	fmt.Printf("%x\n", h.Sum(nil))
}

//\MD4\md4.go
