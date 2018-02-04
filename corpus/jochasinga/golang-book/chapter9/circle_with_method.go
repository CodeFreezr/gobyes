package main

import "fmt"
import "math"

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	// create a circle with coordinate (0,0) and r = 5
	c := Circle{0, 0, 5}
	fmt.Println(c.area())
}
