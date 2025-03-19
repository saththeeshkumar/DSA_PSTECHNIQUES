package main

import "log"

func main() {
	log.Println(Remove_Duplicates_from_Sorted_Array([]int{1, 1, 1, 3, 3, 5, 5}))
	log.Println(Remove_Duplicates_from_Sorted_Array([]int{1, 1, 1}))
	log.Println(Remove_Duplicates_from_Sorted_Array([]int{0}))
	log.Println(Remove_Duplicates_from_Sorted_Array([]int{5, 5, 10, 10, 20}))
	log.Println(Remove_Duplicates_from_Sorted_Array([]int{10, 20, 30, 40, 50}))
}

// Without extra memory(array) and return new length
func Remove_Duplicates_from_Sorted_Array(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}

	i, j := 0, 0
	for ; j < length-1; j++ {
		if nums[j] != nums[j+1] {
			nums[i+1] = nums[j+1]
			i++
		}
	}
	return i + 1
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}
