package eg

// An example interface in a package
type Example interface {
	// Returns the name of this type
	Name() string
	// Unique identifier for the type
	id() uint32
}

// Creates a new value implementing 
// the Example interface
func NewExample() Example {
	return new(concreteType)
}
