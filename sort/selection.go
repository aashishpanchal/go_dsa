/*
Confusing points
Why i < n-1? Because once n-1 elements are sorted, the last one is automatically in place.
Why j = i + 1? Because we search for the minimum element in the remaining unsorted portion of the array.
Why swap arr[min], arr[i]? To move the found minimum element to its correct position in the sorted sequence.
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
