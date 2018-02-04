package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5, 6)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	slice5 := make([]int, 10)
	slice5 = append(slice5, 9, 8, 7)
	fmt.Println("slice5: ", slice5)

	slice6 := make([]string, 3)
	slice6 = append(slice6, "last")
	slice6[0] = "first"
	slice6[1] = "mid"
	slice6[2] = "before_last"
	fmt.Println("slice6: ", slice6)

	slice7 := make([]string, 3)
	copy(slice7, slice6)
	fmt.Println("slice7: ", slice7)
}
