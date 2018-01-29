package structures

import "fmt"

// Visit will execute on each node traversed
type Visit func(n *Node)

// Comparable imposes a total ordering on all types which implement it. This
// ordering is referred to as the type's natural ordering.
type Comparable interface {
	// CompareTo compares this type with the specified type for order. Returns a -1, 0, or 1 if this type
	// is less than, equal to, or greater than the specified type.
	CompareTo(o Comparable) int
}

// Node is an element on the tree with two children
type Node struct {
	Left  *Node
	Right *Node
	Key   Comparable
}

// NewNode creates a new node object from the given key
func NewNode(key Comparable) *Node {
	return &Node{Key: key}
}

func (n *Node) String() string {
	if n.Key == nil {
		return "[]"
	}
	s := "["
	if n.Left != nil {
		s += n.Left.String() + " "
	}
	s += fmt.Sprintf("%v", n.Key)
	if n.Right != nil {
		s += " " + n.Right.String()
	}
	s += "]"
	return s
}

func (n *Node) insert(node *Node) *Node {
	if n == nil {
		return node
	}
	if n.Key.CompareTo(node.Key) <= 0 {
		n.Right = n.Right.insert(node)
	} else {
		n.Left = n.Left.insert(node)
	}
	return n
}

func (n *Node) TraversePre(visit Visit) {
	if n == nil {
		return
	}
	visit(n)
	n.Left.TraversePre(visit)
	n.Right.TraversePre(visit)
}

func (n *Node) TraverseIn(visit Visit) {
	if n == nil {
		return
	}
	n.Left.TraverseIn(visit)
	visit(n)
	n.Right.TraverseIn(visit)
}

func (n *Node) TraversePost(visit Visit) {
	if n == nil {
		return
	}
	n.Left.TraversePost(visit)
	n.Right.TraversePost(visit)
	visit(n)
}
