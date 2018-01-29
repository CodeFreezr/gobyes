package main
import "fmt"

func dup3(in <-chan int) (<-chan int, <-chan int, <-chan int) {
	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			x := <-in
			a <- x
			b <- x
			c <- x
		}
	}()
	return a, b, c
}

func fib() <-chan int {
	x := make(chan int, 2)
	a, b, out := dup3(x)
	go func() {
		x <- 0
		x <- 1
		<-a
		for {
			x <- <-a+<-b
		}
	}()
	return out
}

func main() {
	x := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
}

// See sdh33b.blogspot.com/2009/12/fibonacci-in-go.html
