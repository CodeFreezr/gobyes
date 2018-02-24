package main

import (
	"fmt"
	"strings"
	"unicode"
)

var simple = `
    simple   `

func main() {
	show("original", simple)
	show("leading ws removed", strings.TrimLeftFunc(simple, unicode.IsSpace))
	show("trailing ws removed", strings.TrimRightFunc(simple, unicode.IsSpace))
	// equivalent to strings.TrimFunc(simple, unicode.IsSpace)
	show("both removed", strings.TrimSpace(simple))
}

func show(label, str string) {
	fmt.Printf("%s: |%s| %v\n", label, str, []rune(str))
}

//\Strip-whitespace-from-a-string-Top-and-tail\strip-whitespace-from-a-string-top-and-tail.go
