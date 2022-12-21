package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnTotalWhenGivenTwoNumbers(t *testing.T) {
	result := Add(1, 2)

	assert.Equal(t, 3, result)
}

func TestShouldReturnDifferenceWhenGivenTwoNumbers(t *testing.T) {
	result := Subtract(1, 2)

	assert.Equal(t, -1, result)
}

func TestShouldReturnErrorWhenDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("Expect division by zero error")
	}
}

func TestShouldReturnSineValueForTheOppositeAndHypotenuse(t *testing.T) {
	opposite := 3.0
	hypotenuse := 5.0
	result, _ := Sin(opposite, hypotenuse)

	assert.Equal(t, 0.6, result)
}

func TestShouldReturnCosineValueForGivenAdjacentAndHypotenuse(t *testing.T) {
	adjacent := 3.0
	hypotenuse := 5.0
	result, _ := Cos(adjacent, hypotenuse)

	assert.Equal(t, 0.6, result)
}

func TestShouldReturnTangentValueForGivenOppositeAndAdjacent(t *testing.T) {
	opposite := 3.0
	adjacent := 5.0
	result, _ := Tan(opposite, adjacent)

	assert.Equal(t, 0.6, result)
}

func TestShouldReturnSquareRootForAGivenNumber(t *testing.T) {
	result := Sqrt(18.0)

	assert.Equal(t, 6.0, result)
}
