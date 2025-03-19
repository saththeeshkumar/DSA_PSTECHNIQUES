package main

import (
	"fmt"
	"log"
)

func main() {
	arr := []int{3, 0, 6, 1, 5}
	ans := mergeSort(arr)

	fmt.Println("Sorted array ", ans)
}

func mergeSort(arr []int) []int {

	if len(arr) <= 1 {
		log.Println("returning arr ", arr)
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return Merge(left, right)
}

func Merge(left []int, right []int) []int {
	log.Println("left ", left)
	ans := []int{}

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			ans = append(ans, left[i])
			i++
		} else {
			ans = append(ans, right[j])
			j++
		}
	}
	fmt.Println("After loop ", ans)
	if i < len(left) {
		ans = append(ans, left[i:]...)
	}
	if j < len(right) {
		ans = append(ans, right[j:]...)
	}
	return ans
}
