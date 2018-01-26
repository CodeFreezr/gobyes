package main
import "math/big"
import "fmt"
import "reflect"
import "os"
import "bufio"


func findMethod(name string, value interface{}) (reflect.Method, bool) {
	t := reflect.TypeOf(value)
	for i:=0 ; i<t.NumMethod() ; i++ {
		if m := t.Method(i); m.Name == name {
			return m, true
		}
	}
	return reflect.Method{}, false
}

func main() {
	stack := make([]*big.Int, 0, 100)
	lineReader := bufio.NewReader(os.Stdin)
	for {
		l, _, e := lineReader.ReadLine()
		if e != nil { break }
		s := string(l)
		z := big.NewInt(0)
		z, isInt := z.SetString(s, 0)
		if isInt {
			stack = append(stack, z)
		} else {
			m, ok := findMethod(s, z)
			if ok {
				argc := m.Func.Type().NumIn()
				last := len(stack)-1
				argv := make([]reflect.Value, argc)
				argv[0] = reflect.ValueOf(stack[last])
				for i:=0 ; i<argc-1 ; i++ {
					argv[i+1] =reflect.ValueOf(stack[last-i])
				}
				m.Func.Call(argv)
			}
		}
		fmt.Printf("%v\n", stack)
	}
}
