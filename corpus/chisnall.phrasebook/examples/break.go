package main
import "fmt"

func main() {
	for i := 0 ; i<10 ; i++ {
L:
		for {
			for {
				break L
			}
		}
		fmt.Printf("%d\n", i)
	}
}
