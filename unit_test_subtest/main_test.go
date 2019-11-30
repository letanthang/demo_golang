package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	Input1   int
	Input2   int
	Expected int
}

func TestSum(t *testing.T) {

	tests := []TestCase{{"1+1", 1, 1, 2}, {"1+5", 1, 5, 6}, {"2+5", 2, 5, 7}}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			result := Sum(tc.Input1, tc.Input2)
			assert.Equal(t, tc.Expected, result, "they should be equal")
		})
	}
}

func TestMultiple(t *testing.T) {

	tests := []TestCase{{"1x1", 1, 1, 1}, {"1x5", 1, 5, 5}, {"2x5", 2, 5, 10}}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			result := Multiple(tc.Input1, tc.Input2)
			assert.Equal(t, tc.Expected, result, "they should be equal")
		})
	}

}
