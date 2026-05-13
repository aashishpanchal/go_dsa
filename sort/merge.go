/*
Complexity
Best Case: O(n log n)
Average Case: O(n log n)
Worst Case: O(n log n)
Space Complexity: O(n) - Due to temporary sub-arrays.

Use-cases
- Large datasets: Guaranteed O(n log n) performance.
- Stable sorting: When you need to maintain the relative order of equal elements.
- External sorting: Sorting data that is too large to fit into RAM.
*/

package sort

import "go_dsa/consts"

// Merge sort algorithm
func Merge[T consts.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := Merge(arr[:mid])
	right := Merge(arr[mid:])

	return merge(left, right)
}

// merge helper function to combine two sorted arrays
func merge[T consts.Ordered](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))
	i, j := 0, 0

	// mere sub-arrays with sorted.
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
