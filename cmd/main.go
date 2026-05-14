package main

import (
	"fmt"
	"go_dsa/sort"
)

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Original array\t", arr)
	// arr = sorting.Bubble(arr)
	// arr = sort.Selection(arr)
	// arr = sort.Merge(arr)
	arr = sort.Quick(arr)
	fmt.Println("Sorted array\t", arr)
}
