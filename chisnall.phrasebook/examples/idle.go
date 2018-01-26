package main
import "fmt"
import "os"
import "time"

func later(deferRunning chan bool, delay time.Duration, f func()) {
	t := time.NewTimer(delay)
	for {
		select {
			case cont := <- deferRunning:
				if cont {
					t = time.NewTimer(delay)
				} else {
					f()
					return
				}
			case <- t.C:
				f()
				t = time.NewTimer(delay)
		}
	}
}

func main() {
	deferRunning := make(chan bool)
	buffer := ""
	go later(deferRunning, 3000000000,
			 func() { fmt.Printf("User entered %s\n", buffer) })
	b := make([]byte, 1)
	for b[0] != '\n' {
		os.Stdin.Read(b)
		deferRunning <- true
		buffer += string(b)
	}
	deferRunning <- false
}
