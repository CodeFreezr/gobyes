package structures

import (
	"fmt"
	"testing"
)

type TestInt int

func (a TestInt) CompareTo(o Comparable) int {
	b := o.(TestInt)
	if a < b {
		return -1
	} else if a == b {
		return 0
	} else {
		return 1
	}
}

type TestString string

func (a TestString) CompareTo(o Comparable) int {
	b := o.(TestString)
	if a < b {
		return -1
	} else if a == b {
		return 0
	} else {
		return 1
	}
}

var testCases = []struct {
	input    TestInt
	expected string
}{
	{0, "[0]"},
	{367, "[367]"},
}

var strings = []TestString{
	"F", "B", "G", "A", "D", "I", "H", "C", "E",
}

func createTree(keys []Comparable) *Tree {
	var tree *Tree

	for idx, key := range keys {
		if idx == 0 {
			tree = NewTree(NewNode(key))
		} else {
			tree.Insert(NewNode(key))
		}
	}
	return tree
}

func TestNode_String(t *testing.T) {
	for _, tc := range testCases {
		n := NewNode(tc.input)
		got := n.String()
		if got != tc.expected {
			t.Fatalf("got string '%s', expected '%s'", got, tc.expected)
		}
	}
}

func TestNode_TraversePre(t *testing.T) {
	comps := make([]Comparable, len(strings))
	for i := range strings {
		comps[i] = strings[i]
	}
	tree := createTree(comps)
	fmt.Println(tree.String())
	c := newCrumbtrail(" ")
	tree.Traverse(c.Visit, PreOrder)
	fmt.Println(c.ToString())
}

func TestNode_TraverseIn(t *testing.T) {
	comps := make([]Comparable, len(strings))
	for i := range strings {
		comps[i] = strings[i]
	}
	tree := createTree(comps)
	c := newCrumbtrail(" ")
	tree.Traverse(c.Visit, InOrder)
	fmt.Println("InOrder => ")
	fmt.Println(c.ToString())
}
