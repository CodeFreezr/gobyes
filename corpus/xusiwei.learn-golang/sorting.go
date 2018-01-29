package main

import "fmt"
import "sort"


func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strs:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted:", s)
}
