package main

import (
	"fmt"
	"log"
)

func main() {
	arr := []int{7, 6, 1, 8, 2, 5}

	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			log.Println("after j ", j, " ", arr)
			j--
		}
		log.Println("after i ", i, " ", arr)

	}
	fmt.Println("Sorted Array ", arr)

}
