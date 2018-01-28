package main
import "net/rpc"
import "net"

type Counter struct {
	count int
}
type Arg struct { Increment int }
type Result struct { Value int }
func (c *Counter) Value(in *Arg, out *Result) error {
	c.count += in.Increment
	out.Value = c.count
	return nil
}

func main() {
	server := rpc.NewServer()
	server.RegisterName("GoCounter", new(Counter))
	l, _ := net.Listen("tcp", ":1234")
	server.Accept(l)
}
