package main

func main() {
	var a1 [100]int
	firstHalf := a1[:50]
	secondHalf := a1[50:]
	middle := a1[25:75]
	all := a1[:]
}
