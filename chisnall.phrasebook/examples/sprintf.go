package main
import "fmt"

func main() {
	str := fmt.Sprintf("%T %#v, %d, %v", main, main, 42, "aubergine")
	fmt.Printf("%s\n", str)
}
