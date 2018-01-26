package eg

type concreteType struct {}
const (
	concreteTypeId uint32 = 0
)

func (s concreteType) Name() string {
	return "Concrete type"
}
func (s concreteType) id() uint32 {
	return concreteTypeId
}
func (s concreteType) isEqual(o Example) bool {
	return concreteTypeId == o.id()
}
