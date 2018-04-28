package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	l        []int
	k        int
	expected int
}{
	{
		[]int{6, 90, 234, 11, 39, 700},
		101,
		5,
	},
	{
		[]int{3, 90, 90, 3, 233, 12, 35, 700},
		101,
		5,
	},
}

func TestCounterPairs1(t *testing.T) {
	for _, test := range testCases {
		result := CounterPairs1(test.l, test.k)
		assert.Equal(t, test.expected, result)
	}
}
