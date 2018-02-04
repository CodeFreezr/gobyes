package main
import "fmt"
import "reflect"

func copyAny(inV, outV reflect.Value) {
	t := inV.Type()
	if t.Kind() != reflect.Struct {
		outV.Set(inV)
	} else {
		for i:=0 ; i<t.NumField() ; i++ {
			copyAny(inV.Field(i), outV.Field(i))
		}
	}
}
func duplicate(in interface{}) interface{} {
	outV := reflect.Indirect(reflect.New(reflect.TypeOf(in)))
	copyAny(reflect.ValueOf(in), outV)
	return outV.Interface()
}
type e struct {
	D float32
	E float64
}
type Example struct {
	A int
	B string
	C e
}
func main() {
	a := Example{1, "string", e{12.3, 4.5}}
	b := duplicate(a)
	c := b
	fmt.Printf("%#v == %#v\n", a, c)
}
