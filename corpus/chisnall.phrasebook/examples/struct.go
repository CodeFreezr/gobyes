package main
import "fmt"

type Example struct {
	Val string
	count int
}
func (e *Example) Log() {
	e.count++
	fmt.Printf("%d %s\n", e.count, e.Val)
}

func main() {
	var e Example
	e.Val = "A Message"
	e.Log()
	e.Log()
}
