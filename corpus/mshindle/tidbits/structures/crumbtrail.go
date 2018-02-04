package structures

import "fmt"

type crumbtrail struct {
	sep    string
	buffer string
}

func newCrumbtrail(sep string) *crumbtrail {
	c := &crumbtrail{sep: sep, buffer: ""}
	return c
}

func (c *crumbtrail) Visit(n *Node) {
	c.buffer += fmt.Sprintf("%s%v", c.sep, n.Key)
}

func (c *crumbtrail) ToString() string {
	return c.buffer[1:]
}
