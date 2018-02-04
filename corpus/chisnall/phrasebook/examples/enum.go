package main
import "fmt"

const (
	Red               = (1<<iota)
	Green             = (1<<iota)
	Blue, ColorMask   = (1<<iota), (1<<(iota+1))-1
)

const (
	i complex128 = complex(0, 1)
)

func main() {
	fmt.Printf("%d %d %d %d\n", Red, Green, Blue, ColorMask)
}
