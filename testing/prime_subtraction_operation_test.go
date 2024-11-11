package testing

import (
	"leetcode-daily-challenges/problems"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeSubTrationOperation(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    []int
		expected bool
	}{
		{[]int{4, 9, 6, 10}, true},
		{[]int{6, 8, 11, 12}, true},
		{[]int{5, 8, 3}, false},
	}

	for _, test := range tests {
		assert.Equal(problems.PrimeSubOperation(test.input), test.expected)
	}
}
