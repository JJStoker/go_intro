package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Finds the smallest value of a series of numbers
func Min(xs []float64) float64 {
	min := 0.0
	if len(xs) == 0 {
		return min
	}
	for _, x := range xs {
		if min == 0.0 || x < min {
			min = x
		}
	}
	return min
}

// Finds the largest value of a series of numbers
func Max(xs []float64) float64 {
	max := 0.0
	if len(xs) == 0 {
		return max
	}
	for _, x := range xs {
		if max == 0.0 || x > max {
			max = x
		}
	}
	return max
}
