package main

import "fmt"

type request struct {
	key int
	value string
	ret chan string
} 

func set(m chan request, key int, value string) string {
	result := make(chan string)
	m <- request{key, value, result}
	return <-result
}

func runMap(c chan request) {
	m := make(map[int] string)
	for {
		req := <- c
		old := m[req.key]
		m[req.key] = req.value
		req.ret <- old
	}
}

func main() {
	m := make(chan request)
	go runMap(m)
	fmt.Printf("Set %s\n", set(m, 1, "foo"))
	fmt.Printf("Set %s\n", set(m, 1, "bar"))
}
