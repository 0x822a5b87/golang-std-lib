package main

type Number interface {
	int64 | float64
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K string, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
