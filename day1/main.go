package main

import (
	"errors"
	"math"
)

func main() {
}

func Add(x int, y int) int {
	return x + y
}

func Divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Divide by zero")
	} else {
		return x / y, nil
	}
}

func Sin(x float64) float64 {
	return math.Sin(x)
}

func Cos(x float64) float64 {
	return math.Cos(x)
}

func Tan(x float64) float64 {
	return math.Tan(x)
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
