package main

import "fmt"
import "time"

func main() {
	now := time.Now()
	fmt.Printf("Today is %s\n", now.Format("Monday"))
	fmt.Printf("The time is %s\n", now.Format(time.Kitchen))
}
