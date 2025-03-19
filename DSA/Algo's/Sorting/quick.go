package main

import (
	"fmt"
)

// quickSort recursively sorts the slice using the Quick Sort algorithm.
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2] // Choose the pivot element.
	left := []int{}
	right := []int{}

	// Partition the elements into left and right slices.
	for i, v := range arr {
		// if i == len(arr)/2 { // Skip the pivot element. if
		// 	continue
		// }
		_ = i
		if v == pivot { // Skip the pivot element. if
			continue
		}
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	// Recursively sort the left and right slices.
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5, 5, 1, 10}
	fmt.Println("Original array:", arr)
	sortedArr := quickSort(arr)
	fmt.Println("Sorted array:  ", sortedArr)
}
