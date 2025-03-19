package main

import (
	"fmt"
)

// MergeSort sorts an array using the merge sort algorithm.
func MergeSort(arr []int) []int {
	// Base case: if the array has 1 or no elements, it's already sorted.
	if len(arr) <= 1 {
		return arr
	}

	// Split the array into two halves.
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Merge the sorted halves.
	return merge(left, right)
}

// merge combines two sorted arrays into one sorted array.
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Merge elements while there are elements in both arrays.
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	fmt.Println("After loop ", result)

	// Append any remaining elements from the left array.
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Append any remaining elements from the right array.
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted Array:", arr)

	sortedArr := MergeSort(arr)
	fmt.Println("Sorted Array:", sortedArr)
}
