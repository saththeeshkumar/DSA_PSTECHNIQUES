package main

import "log"

func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	//              0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,11
	log.Println(trapUsingTwoPointer(height))
	log.Println(trapUsingDynamicProgrmming(height))
}

// using Two pointer
func trapUsingTwoPointer(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	left, right := 0, length-1                       // Initialize left and right pointers at both ends
	leftMax, rightMax := height[left], height[right] // Track the maximum height encountered from both sides.At very first, left and right is the most max for both sides
	totalTrap := 0

	// Traverse the array until the two pointers meet
	for left < right {
		// Determine the smaller boundary (leftMax or rightMax)
		if leftMax < rightMax {
			left++                                    // Move left pointer to the right
			leftMax = max(leftMax, height[left])      // compare it with current building height and Update leftMax if necessary
			totalTrap += max(0, leftMax-height[left]) // Calculate trapped water at the current position
		} else {
			right--                                     // Move right pointer to the left
			rightMax = max(rightMax, height[right])     // compare it with current building height and Update rightMax if necessary
			totalTrap += max(0, rightMax-height[right]) // Calculate trapped water at the current position
		}
	}
	return totalTrap
}

// using Dynamic Programming
func trapUsingDynamicProgrmming(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	// Find leftMax array
	leftMax := make([]int, n, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	log.Println(leftMax)
	// Find rightMax array
	rightMax := make([]int, n, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	log.Println(rightMax)
	// Calculate trapped water
	trappedWater := 0
	for i := 0; i < n; i++ {
		// Find min from leftMax and rightMax. then subract current height from min.
		trappedWater += max(0, min(leftMax[i], rightMax[i])-height[i])
	}
	return trappedWater
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
