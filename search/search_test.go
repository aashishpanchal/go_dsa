package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name     string
	input    []int
	target   int
	expected int
}{
	{"should find element in middle", []int{1, 2, 3, 4, 5}, 3, 2},
	{"should find element at start", []int{1, 2, 3, 4, 5}, 1, 0},
	{"should find element at end", []int{1, 2, 3, 4, 5}, 5, 4},
	{"should return -1 if not found", []int{1, 2, 3, 4, 5}, 10, -1},
	{"should work with negative numbers", []int{-10, -5, 0, 5, 10}, -5, 1},
	{"should handle empty array", []int{}, 1, -1},
}

func TestSearch(t *testing.T) {
	t.Run("Linear Search", func(t *testing.T) {
		for _, tc := range testCases {
			result := Linear(tc.input, tc.target)
			assert.Equal(t, tc.expected, result, tc.name)
		}
	})

	t.Run("Binary Search", func(t *testing.T) {
		for _, tc := range testCases {
			result := Binary(tc.input, tc.target)
			assert.Equal(t, tc.expected, result, tc.name)
		}
	})
}
