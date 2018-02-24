package main

import (
	"fmt"
	"time"
)

func from(t0 time.Time) {
	fmt.Println(time.Now().Sub(t0))
}

func empty() {
	defer from(time.Now())
}

func count() {
	defer from(time.Now())
	for i := 0; i < 1e6; i++ {
	}
}

func main() {
	empty()
	count()
}

//\Time-a-function\time-a-function-4.go
