package main

import "log"

func main() {

	arr := []int{4, 1, 5, 2, 4, 1}
	log.Println("ans ", Smallest_sub_ary_with_given_sum(arr, 7))
}

// OWN SOLVE
func Smallest_sub_ary_with_given_sum(nums []int, givenSum int) int {

	if nums == nil {
		return 0
	}

	minLength, sum, startIndex := 0, 0, 0
	for i := startIndex; i < len(nums); {
		if nums[i] == givenSum {
			return 1
		}
		sum += nums[i]
		log.Printf("startIndex %v and I %v  and sum %v", startIndex, i, sum) // below i implemented commented code. but better solution is uncommented one.
		// if sum >= givenSum {
		for sum >= givenSum {
			length := i - startIndex + 1
			log.Println("lenght ", length)
			if minLength == 0 || minLength > length {
				minLength = length
			}
			log.Println("minLength ", minLength)
			sum -= nums[startIndex]
			// sum = 0
			startIndex++
			// i = startIndex
			// continue
		}
		i++
	}
	return minLength

}
