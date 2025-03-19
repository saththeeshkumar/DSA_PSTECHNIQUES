package main

import "log"

func main() {
	// ratings := []int{1, 0, 2}
	// ratings := []int{1, 2, 2}
	ratings := []int{3, 2, 1}
	//               1  2   1
	log.Println(candy(ratings))
}

func candy(ratings []int) int {
	n, sum := len(ratings), 0
	result := make([]int, len(ratings))

	result[0] = 1
	// Left to right
	for i := 1; i < n; i++ {
		result[i] = 1
		if ratings[i] > ratings[i-1] {
			result[i] = 1 + result[i-1]
		}
	}
	// right to left
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && result[i] <= result[i+1] {
			result[i] = result[i+1] + 1
		}
		sum += result[i]
	}
	// Add last index count in total sum
	sum += result[n-1]
	return sum
}
