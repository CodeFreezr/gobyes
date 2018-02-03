package math

import "testing"

// Create a struct of inputs and outputs
type testpair struct {
	values  []float64
	average float64
}

// tests slice with inputs and outputs
var tests = []testpair{
	{[]float64{}, 0},
}

func TestAverage(t *testing.T) {
	// loop through each pair
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
