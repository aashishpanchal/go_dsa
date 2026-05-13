/*
Complexity
Best Case: O(n log n)
Average Case: O(n log n)
Worst Case: O(n^2) - Occurs when the pivot is consistently the smallest or largest element.
Space Complexity: O(n) - In this simple implementation due to new slice allocations.

Use-cases
- General purpose sorting: Usually faster in-practice than Merge Sort.
- When space is a concern (in-place versions use O(log n) space).
- Systems where average-case performance is the primary metric.
*/

package sort

import "go_dsa/consts"

// Quick sort algorithm.
func Quick[T consts.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]

	left := make([]T, 0)
	mid := make([]T, 0)
	right := make([]T, 0)

	for _, x := range arr {
		if x < pivot {
			left = append(left, x)
		} else if x == pivot {
			mid = append(mid, x)
		} else {
			right = append(right, x)
		}
	}

	left = Quick(left)
	right = Quick(right)

	return append(append(left, mid...), right...)
}
