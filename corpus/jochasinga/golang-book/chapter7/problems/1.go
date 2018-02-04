/* `sum` is a function which takes a slice of numbers and
 * adds them together. What would its function signature 
 * look like in Go?
 */

package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 5, 10}
	fmt.Println(sum(mySlice))
}

func sum(numSlice []int) (sum int) {
	sum = 0
	for v := range numSlice {
		sum += numSlice[v]
	}
	return
}
