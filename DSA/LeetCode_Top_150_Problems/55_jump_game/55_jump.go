package main

import (
	"fmt"
	"log"
)

func canJump(nums []int) bool {
	farthest := 0 // Tracks the farthest index we can reach

	for i := 0; i < len(nums); i++ {
		log.Println("initial ", i)
		if i > farthest { // If we reach a point that is not accessible, return false
			return false
		}
		farthest = max(farthest, i+nums[i]) // Update the farthest index
		log.Println("end ", i)
		if farthest >= len(nums)-1 { // If we can reach the last index, return true
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 1, 1, 0, 1, 0}
	fmt.Println(canJump(nums)) // Output: true
}
