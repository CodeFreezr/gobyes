package dynamic

import (
	"fmt"
	"math"
)

var denoms []int = []int{1, 5, 10, 25, 100}

// Coins finds the minimum number of coins to add up to value
func Coins(val int) {
	// min as calculated for each increment
	mins := make([]int, val+1)
	for i := 1; i <= val; i++ {
		mins[i] = math.MaxInt32
	}
	nd := len(denoms)

	for i := 1; i <= val; i++ {
		for j := 0; j < nd; j++ {
			v := denoms[j]
			if v <= i && mins[i-v]+1 < mins[i] {
				mins[i] = mins[i-v] + 1
			}
		}
	}

	fmt.Println("Coin denominations:", denoms)
	fmt.Printf("The number of coins needed to make %d is %d\n\n", val, mins[val])
	// show the calculating table
	fmt.Printf("Sum | Min Coins\n")
	for i := 0; i <= val; i++ {
		fmt.Printf("%3d | %9d\n", i, mins[i])
	}
}
