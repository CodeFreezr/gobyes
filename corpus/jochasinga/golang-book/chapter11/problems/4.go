package simplemath

// Find average, minimum and maximum value of a float64 slice
func Average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func Min(xs []float64) float64 {
	min := xs[0]
	for v := 1; i <= len(xs); i++ {
		if v < min {
			min = v
		}
	}
}

func Max(xs []float64) float64 {
	max := xs[0]
	for i := 1; i <= len(xs); i++ {
		if i > max {
			max = i
		}
	}
}
