package main

import "log"

func main() {
	log.Println(Remove_All_Occurrences_of_an_Element_from_Array([]int{10, 30, 40, 10, 10}, 10))
	log.Println(Remove_All_Occurrences_of_an_Element_from_Array([]int{30, 30, 30, 30, 30, 30}, 30))
	log.Println(Remove_All_Occurrences_of_an_Element_from_Array([]int{10, 50, 30, 17}, 5))
	log.Println(Remove_All_Occurrences_of_an_Element_from_Array([]int{4, 7, 6, 7, 8, 7, 7}, 7))
	log.Println(Remove_All_Occurrences_of_an_Element_from_Array([]int{0, 0, 0, 0, 1, 0, 0}, 0))
}

func Remove_All_Occurrences_of_an_Element_from_Array(nums []int, target int) int {

	newLength := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != target {
			nums[newLength] = nums[i]
			newLength++
		}
	}
	return newLength
}
