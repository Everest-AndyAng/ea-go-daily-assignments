package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnTotalWhenGivenTwoNumbers(t *testing.T) {
	result := Add(1, 2)

	assert.Equal(t, 3, result)
}

func TestShouldReturnErrorWhenDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("Expect division by zero error")
	}
}

func TestShouldReturnSineValueForAGivenNumber(t *testing.T) {
	result := Sin(1.0)

	assert.Equal(t, 0.8414709848078965, result)
}

func TestShouldReturnCosValueForAGivenNumber(t *testing.T) {
	result := Cos(1.0)

	assert.Equal(t, 0.5403023058681398, result)
}

func TestShouldReturnTanValueForAGivenNumber(t *testing.T) {
	result := Tan(1.0)

	assert.Equal(t, 1.557407724654902, result)
}

func TestShouldReturnSquareRootForAGivenNumber(t *testing.T) {
	result := Sqrt(36.0)

	assert.Equal(t, 6.0, result)
}
