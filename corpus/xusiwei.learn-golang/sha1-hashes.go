package main

import "fmt"
import "crypto/sha1"

func main() {
	s := "sha1 this string"

	h := sha1.New()
	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
