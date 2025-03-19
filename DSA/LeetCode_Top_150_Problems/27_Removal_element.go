package main

import "log"

func main() {

	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	log.Println(removeElement(nums1, 2))
	log.Println("nums1 ", nums1)

}

func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}
	j := len(nums) - 1
	var k int
	for i := 0; i <= j; i++ {
		if nums[i] == val {
			for {
				if j <= i {
					return k
				}
				if nums[j] != val {
					k++
					nums[i] = nums[j]
					j--
					break
				}
				j--
			}
		} else {
			k++
		}

	}

	return k
}
