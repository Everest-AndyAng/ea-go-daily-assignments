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

func Subtract(x int, y int) int {
	return x - y
}

func Divide(x float64, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Divide by zero")
	} else {
		return x / y, nil
	}
}

func Sin(opposite float64, hypotenuse float64) (float64, error) {
	return Divide(opposite, hypotenuse)
}

func Cos(adjacent float64, hypotenuse float64) (float64, error) {
	return Divide(adjacent, hypotenuse)
}

func Tan(opposite float64, adjacent float64) (float64, error) {
	return Divide(opposite, adjacent)
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
