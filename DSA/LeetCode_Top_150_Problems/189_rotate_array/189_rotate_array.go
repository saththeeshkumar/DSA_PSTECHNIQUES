package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	n := len(nums)
	// NOTE : If k is greater than n, rotating the array k times is the same as rotating it k % n times.
	//        A full rotation of n times brings the array back to its original state.
	//        So, rotating by k is equivalent to rotating by k % n.
	k = k % n // Handle cases where k > n.
	fmt.Println("k ", k)
	fmt.Println("Original Array:", nums)
	if k == 0 {
		return
	}

	reverse(nums, 0, n-1) // Step 1: Reverse entire array
	fmt.Println("After reversing entire array:", nums)

	reverse(nums, 0, k-1) // Step 2: Reverse first k elements
	fmt.Println("After reversing first", k, "elements:", nums)

	reverse(nums, k, n-1) // Step 3: Reverse remaining elements
	fmt.Println("After reversing last", n-k, "elements:", nums)
}

// Helper function to reverse part of the array
func reverse(nums []int, start, end int) {
	fmt.Printf("Reversing from index %d to %d\n", start, end)
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
		fmt.Println("Intermediate state:", nums)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 9
	rotate(nums, k)
	fmt.Println("Final Rotated Array:", nums)
}

// package main

// import "log"

// func main() {
// 	nums := []int{1, 2, 3, 4, 5, 6, 7}
// 	k := 3
// 	rotate(nums, k)
// 	log.Println("ans ", nums)
// }

// func rotate(nums []int, k int) {
// 	log.Println("nums ", nums)
// 	start := 0
// 	for i := 0; i < len(nums); i++ {
// 		nextPosition := start + k
// 		if nextPosition >= len(nums) {
// 			nextPosition = nextPosition - len(nums)
// 		}

// 		var bkp int
// 		if i == 0 {
// 			bkp = nums[nextPosition]
// 			nums[nextPosition] = nums[i]
// 		} else {
// 			n := nums[nextPosition]
// 			nums[nextPosition] = bkp
// 			bkp = n
// 		}
// 		start = nextPosition
// 		if start < i {
// 			start++
// 		}
// 		log.Println("i and nums ", i, " ", nums)
// 		i++
// 	}

// }
