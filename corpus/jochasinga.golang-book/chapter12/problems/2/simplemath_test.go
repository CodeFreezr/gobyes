package simplemath

import "testing"

// Create a struct of inputs and outputs
type testpair struct {
	values []float64
	min    float64
	max    float64
}

// tests slice with inputs and outputs
var tests = []testpair{
	{[]float64{1, 2, 3, 4, 5}, 1, 5},
	{[]float64{1, 10, 34, 2, 8}, 1, 34},
	{[]float64{1.5, 4.2, 67, 32, 0.2}, 0.2, 67},
	{[]float64{4500, 10e+7, 12, 54.8}, 12, 10e+7},
}

func TestMin(t *testing.T) {
	// loop through each pair
	for _, pair := range tests {
		v := Min(pair.values)
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}
func TestMax(t *testing.T) {
	// loop through each pair
	for _, pair := range tests {
		v := Max(pair.values)
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}
