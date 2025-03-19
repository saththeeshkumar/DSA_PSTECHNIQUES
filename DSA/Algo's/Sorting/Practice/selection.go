package main

import "log"

func main() {
	arr := []int{7, 6, 1, 8, 2, 5}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	log.Println("sorted array ", arr)
}
