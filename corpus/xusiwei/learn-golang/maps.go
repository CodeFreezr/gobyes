package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Println("m:", m)

	v1 := m["one"]
	fmt.Println("m[\"one\"]:", v1)

	fmt.Println("len(m):", len(m))

	delete(m, "one")
	fmt.Println("after delete, m:", m)

	// second return value indicates:
	//  the key was present in map or not
	_, in := m["one"]
	fmt.Println("in:", in)

	n := map[string]int {"foo": 1, "bar": 2}
	fmt.Println("n:", n)
}
