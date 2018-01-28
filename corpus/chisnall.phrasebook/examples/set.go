package main
import "fmt"

func main() {
	set := make(map[string] bool)
	set["A"] = true
	fmt.Printf("%t %t\n", set["A"], set["B"])
	// Setting to false does not remove the value
	set["A"] = false
	for k, v := range set {
		fmt.Printf("%s %t\n", k, v)
	}
	// Remove the element from the set
	delete(set,"A")
	for k, v := range set {
		fmt.Printf("%s %t\n", k, v)
	}
}
