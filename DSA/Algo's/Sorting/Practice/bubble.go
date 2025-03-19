package main

import "log"

func main() {
	arr := []int{7, 6, 1, 8, 2, 5}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	log.Println("sorted array ", arr)

}
