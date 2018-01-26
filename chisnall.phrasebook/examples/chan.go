package main
import "fmt"

func main() {
	m := make(map[int] string)
	m[2] = "First Value"
	c := make(chan bool, 1)
	go func() {
		m[2] = "Second Value"
		c <- true
	}()
	_ = <- c
	fmt.Printf("%s\n", m[2])
}
