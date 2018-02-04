package main

import "fmt"

func main() {

	nums := []int {2, 3, 4}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("sum:", sum)

	// range on array/slice
	// provide both index, value pair
	for i, n := range nums {
		fmt.Println(i, "=>", n)
	}

	// range on map iterates key/value pairs
	m := map[string]string {"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// iterates index, utf8 char value
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
