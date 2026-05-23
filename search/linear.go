package search

import "go_dsa/consts"

// Linear search algorithm.
// Returns the index of the target if found, otherwise -1.
// Complexity: O(n)
func Linear[T consts.Ordered](arr []T, target T) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
