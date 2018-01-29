package main

import "fmt"

func main() {
	var a[5] int
	fmt.Println("a:", a) // by default, an array is *zero-valued*

	a[4] = 100 // access by operator[]
	fmt.Println("after set, a:", a)
	fmt.Println("a[4]:", a[4])

	fmt.Println("len(a):", len(a)) // keyword len

	b := [5] int {1,2,3,4,5} // declare and initlize
	fmt.Println("b:", b)

	var tmp[2][3] int // two dimession
	fmt.Println("len(tmp):", len(tmp))
	fmt.Println("len(tmp[0]):", len(tmp[0]))
	fmt.Println("tmp:", tmp)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			tmp[i][j] = i + j
		}
	}
	fmt.Println("tmp:", tmp)
}
