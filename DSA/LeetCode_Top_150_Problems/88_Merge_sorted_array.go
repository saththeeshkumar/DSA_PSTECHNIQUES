package main

import "log"

func main() {

	nums1 := []int{4, 5, 6, 0, 0, 0}
	m := 3
	nums2 := []int{1, 2, 3}
	n := 3
	merge(nums1, m, nums2, n)
	log.Println("nums1 ", nums1)
}

// func merge(nums1 []int, m int, nums2 []int, n int) {

// 	mergedAry := make([]int, len(nums1), len(nums1))
// 	log.Println("len(mergedAry) ", len(mergedAry))
// 	log.Println("Cap(mergedAry) ", cap(mergedAry))
// 	if m == 0 && n != 0 {
// 		for k, v := range nums2 {
// 			nums1[k] = v
// 		}
// 	}

// 	if n == 0 {
// 		return
// 	}

// 	i, j, k := 0, 0, 0

// 	for i != m && j != n {
// 		if nums1[i] <= nums2[j] {
// 			mergedAry[k] = nums1[i]
// 			i++
// 		} else {
// 			mergedAry[k] = nums2[j]
// 			j++
// 		}
// 		k++
// 		log.Println("i ", i)
// 		log.Println("j ", j)
// 		log.Println("k ", k)
// 	}
// 	log.Println("mergedAry ", mergedAry)

// 	if i == m {
// 		for j < n {
// 			mergedAry[k] = nums2[j]
// 			j++
// 			k++
// 		}
// 	}

// 	if j == n {
// 		for i < m {
// 			mergedAry[k] = nums1[i]
// 			i++
// 			k++
// 		}
// 	}
// 	nums1 = mergedAry
// 	log.Println("mergedAry ", mergedAry)

// }

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	i, j, k := m-1, n-1, m+n-1 // Pointers for nums1, nums2, and merged position

// 	// Merge from the end
// 	for i >= 0 && j >= 0 {
// 		if nums1[i] > nums2[j] {
// 			nums1[k] = nums1[i]
// 			i--
// 		} else {
// 			nums1[k] = nums2[j]
// 			j--
// 		}
// 		k--
// 	}

// 	// If there are remaining elements in nums2, copy them
// 	for j >= 0 {
// 		nums1[k] = nums2[j]
// 		j--
// 		k--
// 	}
// }

func merge(nums1 []int, m int, nums2 []int, n int) {

	m, n, k := m-1, n-1, m+n-1

	for m >= 0 && n >= 0 {

		if nums1[m] >= nums2[n] {
			nums1[k] = nums1[m]
			m--
		} else {
			nums1[k] = nums2[n]
			n--
		}
		k--
	}

	log.Println("num1 ", nums1)
	log.Println("i and k ", m, n, k)

	for n >= 0 {
		nums1[k] = nums2[n]
		n--
		k--
	}

}
