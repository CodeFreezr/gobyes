package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Println("write", i, " as ")
	switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
	}

	// multi-value on a case
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("it's the weekend")
		default:
			fmt.Println("it's a weekday")
	}

	t := time.Now()
	switch {
		case t.Hour() < 12: // variable on case
			fmt.Println("it's before noon")
		default:
			fmt.Println("it's after noon")
	}

	var m int = t.Hour()
	switch t.Hour() {
		case m: // variable on case
			fmt.Println("current hour:", m)
	}

	s := "string"
	switch s {
		case "string": // string on case
			fmt.Println("string as case")
	}
}
