package toy

import (
	"fmt"
	"math/rand"
	"time"
)

// Result from a search
type Result string

// Search is a pre-defined function that takes
// a string as a query and returns a Result
type Search func(query string) Result

// Web search query
var Web = fakeSearch("web")

// Image search query
var Image = fakeSearch("image")

// Video search query
var Video = fakeSearch("video")

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// Google10 runs a collection of searches
func Google10(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

// Google20 runs a collection of searches
func Google20(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		results = append(results, <-c)
	}
	return
}
