package main

import "strings"
import "fmt"

func main() {
	str := "\tThis is a string \n"
	str = strings.Trim(str, " \t\n\r")
	words := strings.Split(str, " ")
	for _, word := range words {
		fmt.Printf("%s\n", word)
	}
}
