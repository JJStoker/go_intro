package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Finds the smallest value of a series of numbers
func Min(xs []float64) float64 {
	var min float64
	if len(xs) == 0 {
		return 0.0
	}
	for i, x := range xs {
		if i == 0 || x < min {
			min = x
		}
	}
	return min
}

// Finds the largest value of a series of numbers
func Max(xs []float64) float64 {
	var max float64
	if len(xs) == 0 {
		return 0.0
	}
	for i, x := range xs {
		if i == 0 || x > max {
			max = x
		}
	}
	return max
}
