package main

import "log"

func main() {
	log.Println(pairWithTargetSumFromSorted([]int{1, 5, 10, 20, 80}, 90))
	log.Println(pairWithTargetSumFromSorted([]int{1, 5, 10, 20, 80}, 40))
	log.Println(pairWithTargetSumFromSorted([]int{-2, -1, 5, 13, 17, 25, 40}, -2))
	log.Println(pairWithTargetSumFromSorted([]int{-2, -1, 5, 13, 17, 25, 40}, 100))
	log.Println(pairWithTargetSumFromSorted([]int{-2, -1, 5, 13, 17, 25, 40}, 39))
}

func pairWithTargetSumFromSorted(nums []int, target int) int {

	if nums == nil {
		return 0
	}

	start, end := 0, len(nums)-1
	for start < end {
		sum := nums[start] + nums[end]
		if sum == target {
			return 1
		}
		if sum < target {
			start++
		} else {
			end--
		}

	}

	return 0

}
