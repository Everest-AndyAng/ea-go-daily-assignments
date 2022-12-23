package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoQuote(t *testing.T) {
	result := GoQuote()
	assert.NotEmpty(t, result)
}
