package eg
import "testing"

func TestExample(t *testing.T) {
	e := NewExample()
    // Don't do this!
	_ = e.(*concreteType)
	if e.Name() != "Concrete type" {
		t.Fail()
	}
	t.Errorf("This test is buggy")
}
