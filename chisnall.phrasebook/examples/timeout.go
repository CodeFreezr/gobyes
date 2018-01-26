package main
import "fmt"
import "time"

func readString(reply chan string) {
	fmt.Printf("Enter some text:\n")
	var s string
	fmt.Scanf("%s", &s)
	reply <- s
}

func timeout(t chan bool) {
	time.Sleep(5000000000)
	t <- true
}

func main() {
	t := make(chan bool)
	s := make(chan string)
	go readString(s)
	go timeout(t)
	select {
		case msg := <- s:
			fmt.Printf("Received: %s\n", msg)
		case <- t:
			fmt.Printf("Timed out\n")
	}
}
