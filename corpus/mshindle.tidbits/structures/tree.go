package structures

// TraversalType defines which type of traversal should occur on a tree
type TraversalType int

const (
	// PreOrder traversal
	PreOrder TraversalType = 1 << iota
	// InOrder traversal
	InOrder TraversalType = 1 << iota
	// PostOrder traversal
	PostOrder TraversalType = 1 << iota
)

// Tree is a binary tree
type Tree struct {
	Root *Node
}

// NewTree creates a new tree object from the given root
func NewTree(root *Node) *Tree {
	return &Tree{Root: root}
}

// String creates a string representation of the tree
func (tr *Tree) String() string {
	return tr.Root.String()
}

// Insert inserts a Node to a Tree without replacement.
func (tr *Tree) Insert(nd *Node) {
	if tr.Root == nd {
		return
	}
	tr.Root = tr.Root.insert(nd)
}

// Traverse moves across the tree executing visit on each node as determined by TraversalType
func (tr *Tree) Traverse(visit Visit, t TraversalType) {
	switch {
	case t&PreOrder == PreOrder:
		tr.Root.TraversePre(visit)
	case t&InOrder == InOrder:
		tr.Root.TraverseIn(visit)
	case t&PostOrder == InOrder:
		tr.Root.TraversePost(visit)
	}

}
