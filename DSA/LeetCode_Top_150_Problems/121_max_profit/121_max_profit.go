package main

import "log"

func main() {

	prices := []int{2, 4, 1}
	log.Println("max profit ", maxProfit(prices))
}

// func maxProfit(prices []int) int {

// 	buyDay, sellDay := -1, -1
// 	for i := 1; i < len(prices); i++ {

// 		if prices[i] < prices[i-1] {
// 			buyDay = i
// 			if sellDay != -1 && prices[i-1] > prices[sellDay] {
// 				sellDay = i - 1
// 			} else if sellDay == -1 {
// 				sellDay = i - 1
// 			}

// 		} else {
// 			buyDay = i - 1
// 			if sellDay != -1 && prices[i] > prices[sellDay] {
// 				sellDay = i
// 			} else if sellDay == -1 {
// 				sellDay = i
// 			}
// 		}

// 		if sellDay <= buyDay {
// 			sellDay = -1
// 		}
// 	}
// 	log.Println("sellday and buyday ", sellDay, buyDay)

// 	if sellDay == -1 || buyDay == -1 {
// 		return 0
// 	}
// 	return prices[sellDay] - prices[buyDay]

// }

// prices := []int{7, 2, 5, 3, 6, 4}
func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
