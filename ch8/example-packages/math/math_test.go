package math

import "testing"

type testpair struct {
	values []float64
	result float64
}

var averageTests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range averageTests {
		v := Average(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

var minTests = []testpair{
	{[]float64{1, 2}, 1},
	{[]float64{-1, 0}, -1},
	{[]float64{2, 3, 4, 1, 5}, 1},
	{[]float64{}, 0},
}

func TestMin(t *testing.T) {
	for _, pair := range minTests {
		v := Min(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

var maxTests = []testpair{
	{[]float64{1, 2}, 2},
	{[]float64{-1, 0}, 0},
	{[]float64{2, 3, 4, 5, 1}, 5},
	{[]float64{}, 0},
}

func TestMax(t *testing.T) {
	for _, pair := range maxTests {
		v := Max(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

/* Exercises
1. See math.go::Average, if len input == 0 -> return 0
2.

*/
