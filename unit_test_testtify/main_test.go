package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := Sum(1, 1)
	assert.Equal(t,2, result, "they should be equal")
	// assert.NotEqual(), assert.Contains()
}

func TestMultiple(t *testing.T) {
	result := Multiple(1, 1)
	assert.Equal(t,1, result, "they should be equal")
}
