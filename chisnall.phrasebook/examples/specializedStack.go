package main
import "fmt"

type stackEntry interface {
	pop() (interface{}, stackEntry)
}
type genericStackEntry struct {
	next stackEntry
	value interface{}
}
func (g genericStackEntry) pop() (interface{}, stackEntry) {
	return g.value, g.next
}
type integerStackEntry struct {
	value int64
	count int
	next stackEntry
}
func (i *integerStackEntry) pop() (interface{}, stackEntry) {
	if (i.count > 0) {
		i.count--
		return i.value, i
	}
	return i.value, i.next
}
type stack struct {
	top stackEntry
	isInteger bool
}

func (s *stack) pushInt(v int64) {
	if (s.isInteger) {
		top := s.top.(*integerStackEntry)
		if top.value == v {
			top.count++
			return
		}
	}
	var e integerStackEntry
	e.value = v
	e.next = s.top
	s.top = &e
	s.isInteger = true
}
func (s *stack) Push(v interface{}) {
	switch val := v.(type) {
		case int64: s.pushInt(val)
		case int: s.pushInt(int64(val))
		default:
			var e genericStackEntry
			e.value = v
			e.next = s.top
			s.top = &e
			s.isInteger = false
	}
}

func (s *stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	v, top := s.top.pop()
	s.isInteger = s.top == top
	s.top = top
	return v
}

func main() {
	var s stack
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
