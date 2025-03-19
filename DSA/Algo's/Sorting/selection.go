package main

import (
	"fmt"
)

func main() {
	fmt.Println("Selection sort starting")
	arr := []int{45, 65, 1, 55, 100}
	// if we use <, then give len(arr). if we use <=,give len(arr)-1. it will make sure we iterate untill last index withou index out of boundary errror.
	// but we have used len(arr)-1 while we use <. because, in sorting if you sort all the elements untill last, the last element will get sorted automatically.
	// so we can omit the last element check in first loop. but, we should untill last in inner loop.

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			// if you use >, you will get least to increase. if you use <, you will get increase to least. // this sentence need to be changed.
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println("Sorted Array ", arr)
}
