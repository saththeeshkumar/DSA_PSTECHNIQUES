package main

import "log"

func main() {

	arr1 := []int{2, 2, 1, 1, 1, 2, 2}
	log.Println("result ", majorityElement(arr1))
}

func majorityElement(nums []int) int {

	count, candidate, mid := 0, 0, len(nums)/2

	for _, num := range nums {
		if count > mid {
			return candidate
		}
		if count == 0 {
			candidate = num
		}

		if candidate == num {
			count++
		} else {
			count--
		}
	}
	return candidate

}

// 1 2 3 4 5 6 7 8 9 10
// n n n n n n n 1 4 3
// 1 4 3 n n n n n n n
// 1 n n n 4 n n n n 3
// 1 n 2 n 3 n n n n n
