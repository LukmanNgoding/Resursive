package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSequence(t *testing.T) {
	var testcases = []struct {
		testcase string
		input    []int
		output   int
	}{
		{
			testcase: "Case 1",
			input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			output:   6,
		},
		{
			testcase: "Case 2",
			input:    []int{-2, -5, 6, -2, -3, 1, 5, -6},
			output:   7,
		},
		{
			testcase: "Case 3",
			input:    []int{-2, -3, 4, -1, -2, 1, 5, -3},
			output:   7,
		},
		{
			testcase: "Case 4",
			input:    []int{-2, -5, 6, -2, -3, 1, 6, -6},
			output:   8,
		},
		{
			testcase: "Case 5",
			input:    []int{-2, -5, 6, 2, -3, 1, 6, -6},
			output:   12,
		},
	}
	for _, value := range testcases {
		result := MaxSequence(value.input)
		t.Run(value.testcase, func(t *testing.T) {
			assert.Equal(t, value.output, result)
		})
	}
}
