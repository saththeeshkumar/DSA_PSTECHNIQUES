package main

import "fmt"

func main() {

	arr := []int{10, 25, 3, 50, 20, 20, 3, 50}
	fmt.Println("unsorted array ", arr)
	sortedAry := quickSort(arr)
	fmt.Println("sorted array ", sortedAry)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := len(arr) / 2
	left := []int{}
	right := []int{}

	for i, v := range arr {
		if i == len(arr)/2 {
			continue
		}
		if v <= pivot {
			left = append(left, pivot)
		}
		if v > pivot {
			right = append(right, pivot)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)

}
