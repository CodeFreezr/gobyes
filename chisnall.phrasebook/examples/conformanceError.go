package example

type Public interface {
	Name() string
}

type Example struct {
	name string
}
func (e Example) Nme() string {
	return e.name
}

func NewExample() Example {
	return Example{"No Name"}
}
