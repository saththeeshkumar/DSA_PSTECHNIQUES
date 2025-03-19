package main

import "log"

func main() {
	log.Println(Remove_Duplicates_Atmost_N_time_from_Sorted_Array([]int{12, 12, 12}, 1))
	log.Println(Remove_Duplicates_Atmost_N_time_from_Sorted_Array([]int{5, 5, 5, 6, 6, 6, 8}, 2))
	log.Println(Remove_Duplicates_Atmost_N_time_from_Sorted_Array([]int{8, 8, 8, 11, 11, 15, 15, 15, 20, 20}, 3))
	log.Println(Remove_Duplicates_Atmost_N_time_from_Sorted_Array([]int{-2, -2, 5, 5, 7}, 2))
	log.Println(Remove_Duplicates_Atmost_N_time_from_Sorted_Array([]int{10, 20, 30, 40, 50}, 1))
}

func Remove_Duplicates_Atmost_N_time_from_Sorted_Array(nums []int, n int) int {
	length := len(nums)
	if length <= n {
		return length
	}

	i, j := n, n
	for ; j < length; j++ {
		if nums[i-n] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
