package main

import "fmt"

func main() {

	arr := []int{10, 25, 3, 50, 20, 20, 3, 50}
	fmt.Println("unsorted array ", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("sorted arrayy ", arr)
}

func quickSort(arr []int, start, end int) {
	if start < end {
		pi := partition(arr, start, end)
		quickSort(arr, start, pi-1)
		quickSort(arr, pi+1, end)
	}

}

func partition(arr []int, start, end int) int {

	pivot := end
	pIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			arr[i], arr[pIndex] = arr[pIndex], arr[i]
			pIndex++
		}
	}
	arr[pIndex], arr[end] = arr[end], arr[pIndex]
	return pIndex
}
