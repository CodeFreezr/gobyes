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

func NewExample() Public {
	return Example{"No Name"}
}
func NewExample2() Public {
	e := Example{"No Name"}
	e.(Public)
	return e
}
