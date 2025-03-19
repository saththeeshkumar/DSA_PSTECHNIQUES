package main

import "fmt"

// partition rearranges elements around a pivot and returns the pivot index.
func partition(arr []int, start, end int) int {
	pIndex := start
	pivot := arr[end]

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			// Swap arr[i] and arr[pIndex]
			arr[i], arr[pIndex] = arr[pIndex], arr[i]
			pIndex++
		}
	}

	// Swap arr[end] (pivot) with arr[pIndex]
	arr[end], arr[pIndex] = arr[pIndex], arr[end]
	return pIndex
}

// quickSort performs Quick Sort on the array.
func quickSort(arr []int, start, end int) {
	if start < end {
		pi := partition(arr, start, end)
		quickSort(arr, start, pi-1) // Sort left subarray
		quickSort(arr, pi+1, end)   // Sort right subarray
	}
}

func main() {
	arr := []int{10, 25, 3, 50, 20, 20, 3, 50}

	fmt.Println("Array in sorted order:")
	quickSort(arr, 0, len(arr)-1)

	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
