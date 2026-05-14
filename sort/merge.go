/*
Complexity
Best Case: O(n log n)
Average Case: O(n log n)
Worst Case: O(n log n)
*/

package sort

import "go_dsa/consts"

// Merge sort algorithm
func Merge[T consts.Ordered](arr []T) []T {
	n := len(arr)
	mergeSort(arr, 0, n-1)
	return arr
}

// merge sort helper
func mergeSort[T consts.Ordered](arr []T, start, end int) {
	if start >= end {
		return
	}

	mid := start + (end-start)/2

	mergeSort(arr, start, mid) // left half
	mergeSort(arr, mid+1, end) // right half

	merge(arr, start, mid, end)
}

// merge helper
func merge[T consts.Ordered](arr []T, start, mid, end int) {
	// temp array
	temp := make([]T, 0, end-start+1)
	i, j := start, mid+1

	// mere sub-arrays with sorted.
	for i <= mid && j <= end {
		if arr[i] < arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}

	// append remaining elements
	temp = append(temp, arr[i:mid+1]...)
	temp = append(temp, arr[j:end+1]...)

	// copy back
	for idx, v := range temp {
		arr[idx+start] = v
	}
}
