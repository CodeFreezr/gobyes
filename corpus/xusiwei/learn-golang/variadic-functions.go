package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...) // expand slice to arg list
}
