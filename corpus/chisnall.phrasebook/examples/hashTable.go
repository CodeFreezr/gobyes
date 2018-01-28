package main
import "fmt"

type Hashable interface {
	Hash() int
	IsEqual(Hashable) bool
}
type HashTable struct {
	table map[int] []Hashable
}
func (h HashTable) Find(value Hashable) Hashable{
	if (h.table == nil) { return nil }
	l := h.table[value.Hash()]
	if l == nil { return nil }
	for _, e := range(l) {
		if value.IsEqual(e) {
			return e
		}
	}
	return nil
}
func (h *HashTable) Add(value Hashable) {
	if h.Find(value) != nil { return }
	hash := value.Hash()
	if (h.table == nil) {
		h.table = make(map[int] []Hashable)
	}
	l := h.table[hash]
	h.table[hash] = append(l, value)
}

type str string
func (s str) Hash() int {
	return len(s)
}
func (s str) IsEqual(other Hashable) bool {
	return s == other.(str)
}

func main() {
	var h HashTable
	h.Add(str("Foo"))
	h.Add(str("Foo"))
	h.Add(str("Bar"))
	h.Add(str("Wibble"))
	fmt.Printf("%v %v %v\n", h.Find(str("Foo")), h.Find(str("Bar")), h.Find(str("Wibble")))
}
