package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["year"] = 1999
	x["age"] = 10
	delete(x, "age")
	fmt.Println("x: ", x)

	y := make(map[string]string)
	y["name"] = "George Ong"
	y["age"] = "18"
	fmt.Println("His name is " + y["name"] + ", he is " + y["age"])
}
