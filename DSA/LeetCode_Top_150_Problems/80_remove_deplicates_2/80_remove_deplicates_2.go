package main

import (
	"fmt"
	"log"
)

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicates(nums)
	log.Println("nums ", nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums) // Already valid if 2 or fewer elements
	}
	// nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	//               0, 1  2  3  4  5  6  7  8
	i := 2 // Pointer to insert next valid element
	for j := 2; j < len(nums); j++ {
		fmt.Printf("Before: i=%d, j=%d, nums=%v\n", i, j, nums)
		fmt.Printf("comparing: i=%d, j=%d,\n", nums[j], nums[i-2])
		if nums[j] != nums[i-2] { // Ensure at most 2 occurrences
			nums[i] = nums[j]
			log.Println("nums ", nums)
			i++
			log.Println("i and j ", i, j)
		}
		fmt.Printf("After : i=%d, j=%d, nums=%v\n\n", i, j, nums)
	}
	return i
}
