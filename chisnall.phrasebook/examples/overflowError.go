package main
import "fmt"

func setRange(i, j int, slice []int)  {
	for n:=0 ; i<j ; i++ {
		slice[n] = i
		n++
	}
}

func main() {
	var arr [100]int
	setRange(20, 50, arr[20:])
	setRange(50, 80, arr[:])
	fmt.Printf("Array: %v\n", arr)
}
