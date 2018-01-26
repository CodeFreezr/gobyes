package main
import "old/netchan"
import "fmt"

func main() {
	counter := 0
	ch := make(chan int, 1)
	server := netchan.NewExporter()
	server.Export("Counter", ch, netchan.Send)
	server.Export("foo", make(chan bool, 12), netchan.Send)
	err := server.ListenAndServe("tcp", "localhost:1234")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	for {
		counter++
		ch <- counter
	}
}
