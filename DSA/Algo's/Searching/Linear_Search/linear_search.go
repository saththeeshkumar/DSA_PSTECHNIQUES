package main

import "log"

func main() {
	arr := []int{1, 4, 5, 2, 3, 6, 9, 8, 7}
	toFind := 5
	for k, v := range arr {
		if v == toFind {
			log.Println("value exist at index no ", k)
			return
		}
	}
}
