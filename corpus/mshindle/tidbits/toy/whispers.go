package toy

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}

// Whisper creates a 100,000 go routines which takes an integer from the right hand
// nighbor, adds 1, and passes it to the left hand neighbor.
func Whisper() {
	var right chan int
	const n = 100000

	leftmost := make(chan int)
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) { c <- 0 }(right)
	fmt.Println(<-leftmost)
}
