package main
import "fmt"
func main() {
	var a [2]int
	slc := a[0:1]
	for i := 0 ; i<15 ; i++ {
		fmt.Printf("Element: %d %d\n", i, slc[i])
	}
}
