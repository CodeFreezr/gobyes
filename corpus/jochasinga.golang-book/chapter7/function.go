package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func total(xs []float64) (total float64) {
	total = 0.0
	for _, v := range xs {
		total += v
	}
	return
}

func main() {
	scores := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(scores))
	fmt.Println(total(scores))
}
