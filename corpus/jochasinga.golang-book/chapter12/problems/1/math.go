package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	if len(xs) == 0 {
		return total
	}

	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
