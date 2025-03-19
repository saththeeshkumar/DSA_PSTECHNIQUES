package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Bubble sort starting")
	arr := []int{50, 5, 100, 2, 87}
	log.Println("Unsorted array ", arr)
	// 50 5 43 2 87
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			log.Println("Arr at index j ", j, "  ---> ", arr)
		}
		log.Println("Arr at index i ", i, "  ---> ", arr)
	}

	fmt.Println("Sorted Array ", arr)
}
