package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Insertion sort started")
	arr := []int{40, 10, 50, 30, 20}
	log.Println("Unsorted array ", arr)
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		index := i

		for index > 0 && arr[index-1] > value {
			arr[index] = arr[index-1]
			index--
		}
		arr[index] = value
		log.Println("Arr at index ", i, "  ---> ", arr)
	}
	fmt.Println("Sorted array ", arr)
}
