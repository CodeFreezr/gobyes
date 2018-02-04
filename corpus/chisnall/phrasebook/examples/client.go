package main
import "net/rpc"
import "fmt"

type Arg struct { Increment int }
type Result struct { Value int }

func main() {
	client, _ := rpc.Dial("tcp", ":1234")
	var r Result
	client.Call("GoCounter.Value", &Arg{1}, &r)
	fmt.Printf("%d\n", r.Value)
}
