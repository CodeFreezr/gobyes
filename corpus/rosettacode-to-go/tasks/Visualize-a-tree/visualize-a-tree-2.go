package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Node struct {
	Name     string
	Children []*Node
}

func main() {
	tree := &Node{"root", []*Node{
		{"a", []*Node{
			{"d", nil},
			{"e", []*Node{
				{"f", nil},
			}}}},
		{"b", nil},
		{"c", nil},
	}}
	enc := toml.NewEncoder(os.Stdout)
	enc.Indent = "   "
	err := enc.Encode(tree)
	if err != nil {
		log.Fatal(err)
	}
}

//\Visualize-a-tree\visualize-a-tree-2.go
