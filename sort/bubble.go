/*
Complexity
Best Case: O(n) - When the array is already sorted (thanks to 'swapped' flag).
Average Case: O(n^2)
Worst Case: O(n^2)
Space Complexity: O(1)
*/

package sort

import "go_dsa/consts"

// Bubble sort algorithm.
func Bubble[T consts.Ordered](arr []T) []T {
	n := len(arr)
	var swapped bool

	for i := 0; i < n-1; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return arr
}
