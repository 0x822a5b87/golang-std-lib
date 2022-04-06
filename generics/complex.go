package main

type NumberPackage interface {
	int | float64 | float32
}

func Add[p NumberPackage](left, right p) p {
	return left + right
}
