package main

import "fmt"

type rect struct {
	width, height int
}

// area method has a *receiver type* of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// methods can be defined for either pointer or value receiver types.
func (r rect) perim() int {
	return 2*(r.width + r.height)
}


func main() {
	r := rect{10, 5}
	fmt.Println(r)
	fmt.Println("area:", r.area())
	fmt.Println("prim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("prim:", rp.perim())

	// Go automatically handles conversion between values and pointers for method calls.
	//  you may want to use pointer receiver type to aviod coping on method calls
	//  or to allow the method to mutate the receiving struct.
}
