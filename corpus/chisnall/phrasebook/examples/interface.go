package main
import "fmt"
import "math"

type cartesianPoint struct {
	x, y float64
}
type polarPoint struct {
	r, θ float64
}

func (p cartesianPoint) X() float64 {return p.x }
func (p cartesianPoint) Y() float64 {return p.y }
func (p polarPoint) X() float64 {
	return p.r*math.Cos(p.θ)
}
func (p polarPoint) Y() float64 {
	return p.r*math.Sin(p.θ)
}
func (self cartesianPoint) Print() {
	fmt.Printf("(%f, %f)\n", self.x, self.y)
}
func (self polarPoint) Print() {
	fmt.Printf("(%f, %f°)\n", self.r, self.θ)
}
type Point interface {
	Printer
	X() float64
	Y() float64
}
type Printer interface {
	Print()
}
func MakePoint(x, y float64) Point {
	return cartesianPoint{x,y}
}

func main() {
	var p Printer
	point := MakePoint(1,2)
	p = point
	p.Print()
	// nil may be assigned to any interface type
	p = nil
	// This will cause a runtime panic
	p.Print()
}
