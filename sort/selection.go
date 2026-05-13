/*
Complexity
Best Case: O(n^2) - Always scans the remaining array.
Average Case: O(n^2)
Worst Case: O(n^2)
Space Complexity: O(1)

Use-cases
- Memory-constrained systems: When memory write operations are expensive (it minimizes the number of swaps).
- Small datasets where simplicity is preferred.
*/

package sort

import "go_dsa/consts"

// Selection sort algorithm
func Selection[T consts.Ordered](arr []T) []T {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// Swap Element
		arr[min], arr[i] = arr[i], arr[min]
	}

	return arr
}
