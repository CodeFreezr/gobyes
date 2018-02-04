package main

import "fmt"
import "time"

func main() {
	parsed, _ := time.Parse("2/1/2006", "15/6/1982")
	now := time.Now()
	parsedSeconds := parsed.Unix()
	fmt.Printf("%d seconds difference\n", now.Unix()-parsedSeconds)
	diff := now.Sub(parsed)
	fmt.Printf("%s difference\n", diff.String())
}
