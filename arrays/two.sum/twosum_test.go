package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TableDriven Test Cases
var testCases = []struct{
	input []int
	target int
	expected [2]int
} {
	{
		[]int{2, 7, 11, 15},
		9,
		[2]int{0, 1},
	},
	{
		[]int{2, 7, 11, 15, 10, 10},
		20,
		[2]int{4, 5},
	},
	{
		[]int{2, 7, 11, 10, 15, -10},
		0,
		[2]int{3, 5},
	},
	{
		[]int{2, 7000000, 11, 10, 15, -10, 58, 500, 1000, 8900, 1000000, 6000000},
		13000000,
		[2]int{1, 11},
	},
}

func TestTwoSum(t *testing.T) {
	for _, test := range testCases  {
		result := TwoSum(test.input, test.target)
		assert.Equal(t, test.expected, result)
	}
}

func BenchmarkTwoSum(b *testing.B) {
	b.ResetTimer()

	for i:= 0; i < b.N ; i++ {
		TwoSum(testCases[3].input, testCases[3].target)
	}
}