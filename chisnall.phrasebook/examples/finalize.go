package main
import "runtime"
import "fmt"

type example struct {
	Str string
}
func finalizer(e *example) {
	fmt.Printf("Finalizing %s\n", e.Str)
}
func NewExample() *example {
	e := new(example)
	runtime.SetFinalizer(e, finalizer)
	return e
}
func main() {
	e := NewExample()
	e.Str = "a structure"
	e = NewExample()
	runtime.GC()
}
