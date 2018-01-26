package main
import "fmt"
func main() {
	var a [2]int
	for i := 0 ; i<15 ; i++ {
		fmt.Printf("Element: %d %d\n", i, a[i])
	}
}
