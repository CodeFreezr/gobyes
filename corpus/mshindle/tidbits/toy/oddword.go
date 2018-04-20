package toy

import (
	"fmt"
	"unicode"

	"github.com/mshindle/tidbits/structures"
)

// Oddword prints out the text with each oddword reversed. From the
// odd word problem:
//
// Consider a character set consisting of letters, a space, and a point. Words consist of one or more,
// but at most 20 letters. An input text consists of one or more words separated from each other by one or more
// spaces and terminated by 0 or more spaces followed by a point. Input should be read from, and including, the
// first letter of the first word, up to and including the terminating point. The output text is to be produced
// such that successive words are separated by a single space with the last word being terminated by a single point.
// Odd words are copied in reverse order while even words are merely echoed. For example, the input string
//  : whats the matter with kansas.
// becomes
//  : whats eht matter htiw kansas.
func Oddword(text string) {
	// print our intro text
	fmt.Printf(" Input string: %s\n", text)

	// l becomes our input stream of characters which we need to process
	a := []rune(text)
	l := listFromRunes(a)

	fmt.Print("Output string: ")
	for l.Head() != nil {
		even(l)
		space(l)
		odd(l)
		space(l)
		terminate(l)
	}
}

func listFromRunes(runes []rune) *structures.LinkedList {
	ll := structures.NewLinkedList()
	for i := range runes {
		item := &structures.Item{runes[i], nil, nil}
		ll.Append(item)
	}
	return ll
}

func odd(l *structures.LinkedList) {
	if l.Head() == nil {
		return
	}
	i := l.Shift()
	r := i.Value.(rune)
	if unicode.IsLetter(r) {
		odd(l)
		fmt.Printf("%c", r)
	} else {
		l.Prepend(i)
	}
}

func even(l *structures.LinkedList) {
	if l.Head() == nil {
		return
	}

	i := l.Shift()
	r := i.Value.(rune)
	if unicode.IsLetter(r) {
		fmt.Printf("%c", r)
		even(l)
	} else {
		l.Prepend(i)
	}
}

func space(l *structures.LinkedList) {
	if l.Head() == nil {
		return
	}

	i := l.Shift()
	r := i.Value.(rune)
	if unicode.IsSpace(r) {
		fmt.Print(" ")
	} else {
		l.Prepend(i)
	}
}

func terminate(l *structures.LinkedList) {
	if l.Head() == nil {
		return
	}

	i := l.Shift()
	r := i.Value.(rune)
	if unicode.IsPunct(r) {
		fmt.Print(".\n")
	} else {
		l.Prepend(i)
	}
}
