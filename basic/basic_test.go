package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:     "Empty",
			input:    []int{},
			expected: [][]int{{}},
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: [][]int{{1}},
		},
		{
			name:  "Two elements",
			input: []int{1, 2},
			expected: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			name:  "Three elements",
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 2, 1},
				{3, 1, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Permute(tt.input)
			assert.ElementsMatch(t, tt.expected, actual)
		})
	}
}
