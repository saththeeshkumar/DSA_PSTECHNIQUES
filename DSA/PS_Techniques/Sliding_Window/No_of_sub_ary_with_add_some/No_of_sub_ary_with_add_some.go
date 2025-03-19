package main

import "log"

func main() {

	arr := []int{4, 1, 5, 2, 4, 1}
	log.Println("ans ", numOfSubarrays(arr))
}

// not understood
func numOfSubarrays(arr []int) int {
	evenCount, oddCount := 1, 0
	currentSum, result := 0, 0

	for _, num := range arr {
		log.Println("num ", num)
		currentSum += num
		log.Println("currentSum ", currentSum)
		if currentSum%2 == 0 {
			result += oddCount
			log.Println("result inside Even ", result)
			evenCount++
			log.Println("evenCount ", evenCount)
		} else {
			result += evenCount
			log.Println("result inside Add ", result)
			oddCount++
			log.Println("oddCount ", oddCount)
		}
	}

	return result
}
