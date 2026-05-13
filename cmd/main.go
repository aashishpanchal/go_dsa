package main

import (
	"fmt"
	"go_dsa/sort"
)

func main() {
	arr := []int{5, 1, 2, 4, 3, 0}
	fmt.Println("Original array\t", arr)
	// arr = sorting.Bubble(arr)
	arr = sort.Selection(arr)
	fmt.Println("Sorted array\t", arr)
}
