package main

import "fmt"

func main() {
	// arr := []int{2, 1, 5, 1, 3, 2} // k := 3 // 9
	// arr := []int{1, 5, -1, 6, 3, 2}  // k := 3 // 11
	arr := []int{4, 2, 3, 5, 1} // 8
	k := 2
	fmt.Println("Maximum sum of subarray of size", k, ":", FindMaxSumOfGivenKSubAry(arr, k)) // Output: 9
}

func FindMaxSumOfGivenKSubAry(nums []int, k int) int {
	if len(nums) < k {
		return -1 // Edge case: If k > array length, return -1
	}

	windowSum, maxSum := 0, 0
	for i, v := range nums {
		windowSum += v // Add current element to the window

		// When we reach size k, start sliding the window
		if i >= k-1 {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			windowSum -= nums[i-(k-1)] // Remove first element of the window
		}
	}
	return maxSum

}
