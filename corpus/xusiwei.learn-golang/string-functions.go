package main

import s "strings"
import "fmt"


var p = fmt.Println

func main() {

	p("Contains:   ", s.Contains("test", "est"))
	p("Count:      ", s.Count("test", "t"))
	p("HasPrefix:  ", s.HasPrefix("test", "te"))
	p("HasSuffix:  ", s.HasSuffix("test", "st"))
	p("Index:      ", s.Index("test", "e"))
	p("Join:       ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:     ", s.Repeat("a", 5))
	p("Replace:    ", s.Replace("foo", "o", "0", -1))
	p("Replace:    ", s.Replace("foo", "o", "0", 1))
	p("Split:      ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:    ", s.ToLower("TEST"))
	p("ToUpper:    ", s.ToUpper("test"))
	p()

	p("len:", len("hello"))
	p("Char:", "hello"[1])
}
