package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}

	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity)

	elysium := location{long: 154.7}
	fmt.Println(elysium)
	spirit := location{-14.5684, 175.472636}
	fmt.Println(spirit)
	curiosity := location{-4.5895, 137.4417}

	fmt.Printf("%v\n", curiosity)
	fmt.Printf("%+v\n", curiosity)
}
