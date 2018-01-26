package main
import "fmt"

func main() {
	str := "A string"
	bytes := make([]byte, len(str))
	copy(bytes, str)
	strCopy := string(bytes)
	if strCopy != str { panic("Copying failed!") }
	fmt.Printf("%#v\n", strCopy)
}
