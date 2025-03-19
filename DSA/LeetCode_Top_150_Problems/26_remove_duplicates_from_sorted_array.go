package main

import "log"

func main() {

	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//             0, 1, 2, 3, 4, 5, 6, 7, 8, 9
	log.Println(removeDuplicates(nums1))
	log.Println("nums1 ", nums1)

}

// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	// nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
// 	//                0, 1, 2, 3, 4, 5, 6, 7, 8, 9
// 	//                   i  j        f
// 	i := 0
// 	j := 1
// 	for i < len(nums) {
// 		log.Println("Checking for ", nums[i], nums[j])
// 		log.Println("Checking index ", i, j)
// 		if j < len(nums) && nums[i] == nums[j] {
// 			// f := j + 1
// 			// for {
// 			// 	if f >= len(nums)-1 {
// 			// 		log.Println("returned")
// 			// 		return i + 1
// 			// 	}
// 			// 	if nums[j] != nums[f] {
// 			// 		nums[j] = nums[f]
// 			// 		i = j
// 			// 		j = f + 1
// 			// 		break
// 			// 	}
// 			// 	f++
// 			// }
// 			nums[i+1] =
// 		} else {
// 			j++
// 			i++
// 		}
// 		log.Println("nums of i ", i, nums)
// 	}
// 	return i + 1

// }

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1 // Length of unique elements
}
