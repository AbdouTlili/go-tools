package arrays

//the aberage of elements of an array

func Average(array []float64) float64 {
	// keep the sum here
	sum := 0.0

	for _, value := range array {
		sum += value
	}

	return sum / float64(len(array))
}
