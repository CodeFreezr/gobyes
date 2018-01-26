package main
import "fmt"

type Example struct {
	Val string
	count int
}

type integer int
func (i integer) log() {
	fmt.Printf("%d\n", i);
}
func (e *Example) Log() {
	e.count++
	fmt.Printf("%d %s\n", e.count, e.Val)
}

type logger interface {
	log()
}
func main() {
	var i integer
	var l logger
	i = 42
	l = i
	i.log()
	l.log()
}
