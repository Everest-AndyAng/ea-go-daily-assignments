package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Ideally, we wat to mock external dependencies and mock it's behavior
// So that we don't depend on the behavior/implemntation, we just want to delegate
// In order to do that, we would need to more complex implementation by introducing interface and mocking the depedency
func TestGoQuote(t *testing.T) {
	result := GoQuote()
	assert.NotEmpty(t, result)
}
