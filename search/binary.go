package search

import "go_dsa/consts"

// Binary search algorithm.
// The input array MUST be sorted.
// Returns the index of the target if found, otherwise -1.
// Complexity: O(log n)
func Binary[T consts.Ordered](arr []T, target T) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
