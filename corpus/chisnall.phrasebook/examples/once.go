package main
import "fmt"
import "sync"

type LazyInit struct {
	once sync.Once
	value int
}
func (s *LazyInit) Value() int {
	s.init()
	return s.value
}
func (s *LazyInit) init() {
	s.once.Do(func() { s.value = 42 })
}
func (s *LazyInit) SetValue(v int) {
	s.value = v
}

func main() {
	var l LazyInit
	fmt.Printf("%d\n", l.Value())
	l.SetValue(12)
	fmt.Printf("%d\n", l.Value())
}
