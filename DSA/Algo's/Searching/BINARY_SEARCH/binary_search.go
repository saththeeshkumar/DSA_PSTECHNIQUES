package main

import "log"

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 15, 18}
	toFind := 19
	result, index := binarySearch(arr, toFind)
	log.Println("result and index ", result, " ", index)

}

func binarySearch(arr []int, toBeFind int) (bool, int) {

	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == toBeFind {
			return true, mid
		} else if arr[mid] > toBeFind {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false, -1
}
