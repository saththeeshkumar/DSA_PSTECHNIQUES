package main

import "fmt"

func mergeSort(arr []int, start, end int) {
	if start < end {
		mid := start + (end-start)/2

		// Sort the left half
		mergeSort(arr, start, mid)

		// Sort the right half
		mergeSort(arr, mid+1, end)

		// Merge the two halves
		merge(arr, start, mid, end)
	}
}

func merge(arr []int, start, mid, end int) {
	left, right := start, mid+1

	// Temporary array for merging
	temp := make([]int, 0, end-start+1)

	// Merge both halves into temp
	for left <= mid && right <= end {
		if arr[left] <= arr[right] {
			temp = append(temp, arr[left])
			left++
		} else {
			temp = append(temp, arr[right])
			right++
		}
	}

	// Add remaining elements of the left half
	for left <= mid {
		temp = append(temp, arr[left])
		left++
	}

	// Add remaining elements of the right half
	for right <= end {
		temp = append(temp, arr[right])
		right++
	}

	// Copy the merged elements back to the original array
	for i, val := range temp {
		arr[start+i] = val
	}
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Before sorting:", arr)

	mergeSort(arr, 0, len(arr)-1)

	fmt.Println("After sorting:", arr)
}
