package main
import "fmt"
import "io"
import "reflect"

func checkType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
		case reflect.Int, reflect.Uint,
			 reflect.Int8, reflect.Int16,
			 reflect.Int32, reflect.Int64,
			 reflect.Uint8, reflect.Uint16,
			 reflect.Uint32, reflect.Uint64,
			 reflect.Uintptr:
			 fmt.Printf("%v is some kind of integer\n", v)
		case reflect.Struct:
			 fmt.Printf("%#v is a structure\n", v)
			 r := reflect.TypeOf(struct{ io.Reader}{})
			 r = r.Field(0).Type
			 if t.Implements(r) {
				 fmt.Printf("%#v implements the io.Reader interface\n", v)
			 }
	}
}

type NullReader struct {}
func (_ NullReader) Read(_ []byte) (n int, err error) {
	return 0, nil
}

func main() {
	checkType(12)
	checkType(NullReader{})
}
