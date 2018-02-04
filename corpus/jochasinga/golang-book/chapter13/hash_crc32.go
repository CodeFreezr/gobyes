package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)
}
