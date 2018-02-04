package main
import "fmt"

type stackEntry struct{
	next *stackEntry
	value interface{}
}
type stack struct {
	top *stackEntry
}

func (s *stack) Push(v interface{}) {
	var e stackEntry
	e.value = v
	e.next = s.top
	s.top = &e
}
func (s *stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	v := s.top.value
	s.top = s.top.next
	return v
}

type Stack interface {
	Push(interface{})
	Pop() interface{}
}
func NewStack() Stack {
	return &stack{}
}

func main() {
	s := NewStack()
	s.Push(12)
	s.Push(12)
	s.Push("foo")
	s.Push("bar")
	fmt.Printf("%v\n", s.Pop())
	fmt.Printf("%v\n", s.Pop())
	fmt.Printf("%v\n", s.Pop())
	fmt.Printf("%v\n", s.Pop())
	fmt.Printf("%v\n", s.Pop())
}
