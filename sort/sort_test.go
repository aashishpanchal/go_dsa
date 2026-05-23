package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test cases shared across all sorting algorithms
var testCases = []struct {
	name     string
	input    []int
	expected []int
}{
	{"should sort an already sorted array", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{"should sort a reverse sorted array", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	{"should sort an unsorted array", []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
	{"should handle single element", []int{42}, []int{42}},
	{"should sort two elements", []int{2, 1}, []int{1, 2}},
	{"should handle duplicate values", []int{3, 3, 3, 1, 1, 2}, []int{1, 1, 2, 3, 3, 3}},
	{"should sort negative numbers", []int{-3, 5, -1, 0, 2}, []int{-3, -1, 0, 2, 5}},
	{"should handle all same values", []int{7, 7, 7, 7}, []int{7, 7, 7, 7}},
	{"should sort large range values", []int{100, -100, 50, -50, 0}, []int{-100, -50, 0, 50, 100}},
}

// helper to copy slice (because sort modifies in-place)
func copySlice(arr []int) []int {
	cp := make([]int, len(arr))
	copy(cp, arr)
	return cp
}

func TestSort(t *testing.T) {
	t.Run("Bubble Sort", func(t *testing.T) {
		for _, tc := range testCases {
			result := Bubble(copySlice(tc.input))
			assert.Equal(t, tc.expected, result, tc.name)
		}
	})

	t.Run("Insertion Sort", func(t *testing.T) {
		for _, tc := range testCases {
			result := Insertion(copySlice(tc.input))
			assert.Equal(t, tc.expected, result, tc.name)
		}
	})

	t.Run("Selection Sort", func(t *testing.T) {
		for _, tc := range testCases {
			result := Selection(copySlice(tc.input))
			assert.Equal(t, tc.expected, result, tc.name)
		}
	})

	t.Run("Merge Sort", func(t *testing.T) {
		for _, tc := range testCases {
			result := Merge(copySlice(tc.input))
			assert.Equal(t, tc.expected, result, tc.name)
		}
	})

	t.Run("Quick Sort", func(t *testing.T) {
		for _, tc := range testCases {
			result := Quick(copySlice(tc.input))
			assert.Equal(t, tc.expected, result, tc.name)
		}
	})
}
