package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	fmt.Println(minVal(x))
}

func minVal(arr []int) int {
	min := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
