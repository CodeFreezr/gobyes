package main

import "fmt"

type Person struct {
	name string
}

// People can talk...
func (someguy *Person) Talk() {
	fmt.Println("Hi, my name is", someguy.name)
}

type Android struct {
	Person // anonymous field without a name
	model  string
}

func main() {
	p := new(Person)
	p.name = "Roger"
	p.Talk()

	a := new(Android)
	a.name = "HAL2000"
	// Since androids are persons, androids can talk too!
	a.Talk()
}
