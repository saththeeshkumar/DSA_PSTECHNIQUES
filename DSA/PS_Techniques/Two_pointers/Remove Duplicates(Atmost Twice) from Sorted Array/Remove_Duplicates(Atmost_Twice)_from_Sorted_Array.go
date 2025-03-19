package main

import "log"

func main() {
	log.Println(Remove_Duplicates_Almost_Twice_from_Sorted_Array([]int{12, 12, 12}))
	log.Println(Remove_Duplicates_Almost_Twice_from_Sorted_Array([]int{5, 5, 5, 6, 6, 6, 8}))
	log.Println(Remove_Duplicates_Almost_Twice_from_Sorted_Array([]int{8, 8, 8, 11, 11, 15, 15, 15, 20, 20}))
	log.Println(Remove_Duplicates_Almost_Twice_from_Sorted_Array([]int{-2, -2, 5, 5, 7}))
	log.Println(Remove_Duplicates_Almost_Twice_from_Sorted_Array([]int{10, 20, 30, 40, 50}))
}

func Remove_Duplicates_Almost_Twice_from_Sorted_Array(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}

	i, j := 2, 2
	for ; j < length; j++ {
		if nums[i-2] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
