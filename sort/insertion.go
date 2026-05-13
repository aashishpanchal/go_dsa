/*
Confusing points
Why i=1? Because the first element is already considered sorted.
Why j=i-1? Because we compare the 'key' with elements to its left.
Why arr[j+1]=key? Because after the loop ends, j is either -1 or points to an element smaller than or equal to the key.
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
