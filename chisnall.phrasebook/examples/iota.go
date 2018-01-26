package main
import "fmt"

func main() {
	// iota always expands to 0 in this usage
	fmt.Printf("%d %d\n", iota, iota)
	fmt.Printf("%d %d\n", iota, iota)
}
