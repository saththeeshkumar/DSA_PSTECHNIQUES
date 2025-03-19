package main

import "log"

func main() {
	a := []int{3, 0, 6, 1, 5}
	log.Println(hIndex(a))
}

func hIndex(nums []int) int {

	sortedCitations := MergeSort(nums)
	h := 0
	for i, v := range sortedCitations {
		if v >= i+1 {
			h++
		} else {
			break
		}
	}
	return h
}

func MergeSort(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	return merge(left, right)

}

func merge(left, right []int) []int {

	ans := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			ans = append(ans, left[i])
			i++
		} else {
			ans = append(ans, right[j])
			j++
		}
	}

	if i < len(left) {
		ans = append(ans, left[i:]...)
	}
	if j < len(right) {
		ans = append(ans, right[j:]...)
	}

	return ans

}
