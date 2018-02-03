/* Write a function with one variadic parameter that finds the
 * greatest number in a list of numbers
 */

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(greatest(arr...))
}

func greatest(args ...int) int {
	greatest := args[0]
	for v := 1; v <= len(args); v++ {
		if v > greatest {
			greatest = v
		}
	}
	return greatest
}
