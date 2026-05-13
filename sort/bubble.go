/*
Confusing points
Why i < n-1? Because after n-1 passes, the last element is guaranteed to be in its correct position.
Why j < n-i-1? Because in each pass i, the largest i elements are already bubbled to the end, so we don't need to check them again.
Why 'swapped' flag? It's an optimization to stop early if the array is already sorted.
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
