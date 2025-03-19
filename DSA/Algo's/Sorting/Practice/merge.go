package main

import "log"

func main() {
	arr := [11]int{2, 8, 1, 6, 9, 4, 3, 7, 0, 5, 9}
	log.Println("Unsorted array ", arr)
	ans := mergeSort(arr)
	log.Println("Sorted array ", ans)

}

func mergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	log.Println("left ", left)
	log.Println("right ", right)

	temp := make([]int, 0, len(left)+len(right))
	// temp := []int{}
	var i, j = 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			temp = append(temp, left[i])
			i++
		} else {
			temp = append(temp, right[j])
			j++
		}
	}
	if i < len(left) {
		temp = append(temp, left[i:]...)
	}
	if j < len(right) {
		temp = append(temp, right[j:]...)
	}
	return temp
}
