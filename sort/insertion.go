/*
Complexity
Best Case: O(n) - When the array is already sorted.
Average Case: O(n^2)
Worst Case: O(n^2)
Space Complexity: O(1)
*/

package sort

import "go_dsa/consts"

// Insertion sort algorithm.
func Insertion[T consts.Ordered](arr []T) []T {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		// Go While Loop
		for j >= 0 && arr[j] > key {
			// Shifting
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	return arr
}
