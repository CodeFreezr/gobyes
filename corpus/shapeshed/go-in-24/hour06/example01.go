package main

import "fmt"

func main() {
	var cheeses [2]string
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses de Bourgogne"
	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])
	fmt.Println(cheeses)
}
