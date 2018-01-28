package main
import "fmt"

type Any interface {}

func main() {
	a := make(map[int] string)
	b := make(map[Any] int)
	a[12] = "A string in a map"
	b[12] = 12
	b["12"] = 13
	b[12.0] = 14
	fmt.Printf("%s %d %d\n", a[12], b[12], b["12"])
}
