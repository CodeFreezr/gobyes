package main

import "fmt"
import "time"

func main() {
	for year := 2008; year <= 2121; year++ {
		if time.Date(year, 12, 25, 0, 0, 0, 0, time.UTC).Weekday() ==
			time.Sunday {
			fmt.Printf("25 December %d is Sunday\n", year)
		}
	}
}

//\Day-of-the-week\day-of-the-week.go
