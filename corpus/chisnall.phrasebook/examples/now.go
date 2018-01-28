package main
import "fmt"
import "time"

func main() {
	now := time.Now();
	fmt.Printf("%d seconds since the Epoc\n", now.Unix())
	fmt.Printf("%d nanoseconds since the Epoc\n", now.UnixNano())
}
