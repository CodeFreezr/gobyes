package main

import "os"

func foo(n int) (int) {
	if n < 0 {
		panic("unexpected n")
	}
	return n*n
}

func main() {

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	foo(-1)
}
