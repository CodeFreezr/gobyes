package main
import "math/big"
import "fmt"

func main() {
	var n int
	fmt.Printf("Compute how many Fibonacci numbers? ")
	fmt.Scanf("%d", &n)
	last := big.NewInt(1)
	current := big.NewInt(1)
	for i := 0 ; (i < n) && (i < 2) ; i++ {
		fmt.Printf("1\n")
	}
	for i := 2 ; i < n ; i++ {
		last.Add(last, current)
		tmp := last
		last = current
		current = tmp
		fmt.Printf("%s\n", current.String())
	}
}
