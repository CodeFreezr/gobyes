package main

import (
	"fmt"
	"time"
)

type Animal struct {
	Kingdom string
	Specie  string
	Age     rune // alias for uint32
	Pod     uint64
}

func (a *Animal) Born() {
	fmt.Println("I'm born on", time.Now())
}

func (a *Animal) MakeSound() {
	fmt.Println("...")
}

func (a *Animal) Walk() {
	fmt.Println("I'm walking...")
}

type Dog struct {
	Animal        // Dog "is-a" Animal
	Name   string // Dog "has-a" Name
}

func main() {
	a := new(Animal)
	a.Pod = 4
	a.Born()
	a.MakeSound()
	a.Walk()
	fmt.Println()
	max := new(Dog)
	max.Name = "Max"
	max.Pod = 4
	fmt.Println(max.Pod)
	max.Born()
	max.MakeSound()
	max.Walk()
}
