package main
import "fmt"
import "time"

func main() {
	time.AfterFunc(2000000000, func () {
		fmt.Printf("Timer expired\n")
	})
	timer := time.NewTimer(3000000000)
	time :=  <- timer.C
	fmt.Printf("Current time: %d nanoseconds\n", time.UnixNano())
}
