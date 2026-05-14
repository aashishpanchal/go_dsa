/*
Complexity
Best Case: O(n log n)
Average Case: O(n log n)
Worst Case: O(n^2)
*/

package sort

import "go_dsa/consts"

// Quick sort algorithm.
func Quick[T consts.Ordered](arr []T) []T {
	n := len(arr)
	quickSort(arr, 0, n-1)
	return arr
}

// merge sort helper
func quickSort[T consts.Ordered](arr []T, start, end int) {
	if start >= end {
		return
	}

	pivIdx := partition(arr, start, end)

	quickSort(arr, start, pivIdx-1) // right array
	quickSort(arr, pivIdx+1, end)   // left array
}

func partition[T consts.Ordered](arr []T, start, end int) int {
	idx, pivot := start-1, arr[end]

	for j := start; j < end; j++ {
		if arr[j] <= pivot {
			idx++
			// Swap
			arr[j], arr[idx] = arr[idx], arr[j]
		}
	}

	idx++

	// Swap
	arr[end], arr[idx] = arr[idx], arr[end]

	return idx
}
