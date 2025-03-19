package main

import "log"

func main() {
	// 1 0 1 0 1 0
	log.Println(Move_All_Zeros_to_the_End([]int{1, 0, 1, 0, 1, 0}))
	log.Println(Move_All_Zeros_to_the_End([]int{0, 0, 0, 0, 1}))
	log.Println(Move_All_Zeros_to_the_End([]int{1, 1, 1, 1, 1, 1}))
	log.Println(Move_All_Zeros_to_the_End([]int{0, 0, 1, 1, 0, 0, 1, 1}))
	log.Println(Move_All_Zeros_to_the_End([]int{0, 0, 0, 0, 0, 0}))
}

func Move_All_Zeros_to_the_End(nums []int) []int {

	start, end := 0, len(nums)-1

	for start < end {
		if nums[start] == 0 {
			if nums[end] == 0 {
				end--
				continue
			} else {
				nums[start], nums[end] = nums[end], nums[start]
			}
		}
		start++
	}
	return nums
}
