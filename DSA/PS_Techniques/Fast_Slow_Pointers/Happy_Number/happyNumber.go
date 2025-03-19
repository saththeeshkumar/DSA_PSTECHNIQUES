package main

import "fmt"

//A Happy Number is a number that, when you repeatedly replace it with the sum of the squares of its digits, eventually becomes 1. If it enters a cycle that does not include 1, then it is called an Unhappy (or Sad) Number.

// Notes:
// All numbers that eventually reach 1 are happy.
// Numbers that enter a cycle without reaching 1 are unhappy.
// 1 is a happy number, and any number that reaches 1 remains 1.
func main() {
	numbers := []int{19, 2, 7, 20}
	for _, num := range numbers {
		if isHappy(num) {
			fmt.Printf("%d is a happy number.\n", num)
		} else {
			fmt.Printf("%d is not a happy number.\n", num)
		}
	}
}

// Used two pointers. we can use map also.
func isHappy(num int) bool {
	slow := num
	fast := sumOfSqures(num)

	for fast != 1 && slow != fast {
		slow = sumOfSqures(slow)
		fast = sumOfSqures(sumOfSqures(fast))
	}

	return fast == 1
}

func sumOfSqures(num int) int {
	sum := 0
	for num > 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	return sum

}
