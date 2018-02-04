package simplemath

// Find minimum and maximum value of a float64 slice
func Min(xs []float64) float64 {
	min := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] < min {
			min = xs[i]
		}
	}
	return min
}

func Max(xs []float64) float64 {
	max := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] > max {
			max = xs[i]
		}
	}
	return max
}
