package main

import "log"

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

// Example 2:
// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]

func main() {
	nums := []int{1, 2, 3, 4}
	log.Println(productExceptSelf(nums))

}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	result[0] = 1

	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
		log.Println("suffix ", suffix)
	}
	return result
}

// package main

// import (
// 	"fmt"
// 	"log"
// )

// func productExceptSelf(nums []int) []int {
// 	n := len(nums)
// 	result := make([]int, n)

// 	// Step 1: Compute Prefix Product
// 	result[0] = 1 // There is no element before the first element, so initialize it as 1
// 	for i := 1; i < n; i++ {
// 		result[i] = result[i-1] * nums[i-1]
// 	}
// 	log.Println("result prefix ", result)

// 	// Step 2: Compute Suffix Product while updating result
// 	suffix := 1
// 	for i := n - 1; i >= 0; i-- {
// 		result[i] *= suffix
// 		log.Println("resutl of i ", result[i])
// 		suffix *= nums[i] // Update suffix product
// 		log.Println("suffix ", suffix)
// 	}
// 	log.Println("result suffix ", result)
// 	return result
// }

// func main() {
// 	nums := []int{4, 2, 3, 4}
// 	log.Println("nums ", nums)
// 	fmt.Println(productExceptSelf(nums)) // Output: [24, 12, 8, 6]
// }
