package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	b, err := json.MarshalIndent(tree, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

//\Visualize-a-tree\visualize-a-tree-1.go
