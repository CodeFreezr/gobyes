package main

import (
	"fmt"
)

func main() {
	str := "Mary had a %s lamb"
	txt := "little"
	out := fmt.Sprintf(str, txt)
	fmt.Println(out)
}

//\String-interpolation--included-\string-interpolation--included-.go
