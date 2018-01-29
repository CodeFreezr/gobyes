package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0 // dereference with operator*, same as C/C++
}


func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("after zeroval:", i)

	zeroptr(&i) // get address by operator&, same as C/C++
	fmt.Println("after zeroptr:", i)

	fmt.Println("pointer:", &i)
}
