package main

import "log"

func main() {
	// gas := []int{1, 2, 3, 4, 5}
	// cost := []int{3, 4, 5, 1, 2}

	// gas := []int{2, 3, 4}
	// cost := []int{3, 4, 3}

	gas := []int{5, 1, 2, 3, 4}  //16
	cost := []int{4, 4, 1, 5, 1} //15
	log.Println(canCompleteCircuit(gas, cost))
}

// func canCompleteCircuit(gas []int, cost []int) int {

// 	n := len(gas)
// 	lack, more, startIndex := 0, 0, -1

// 	for i := 0; i < n; i++ {
// 		if startIndex == -1 && gas[i] < cost[i] {
// 			lack += cost[i] - gas[i]
// 			continue
// 		} else if startIndex == -1 {
// 			startIndex = i
// 		}

// 		more += gas[i] - cost[i]
// 		if more < 0 {
// 			return -1
// 		} else if more == 0 {
// 			startIndex = -1
// 		}
// 	}

// 	if more < lack {
// 		return -1
// 	}

// 	return startIndex
// }

// Understanding the Problem
// For a valid circuit:

// 1. Total Gas Available ≥ Total Gas Needed
// ∑gas[i] ≥ ∑cost[i]
// → Otherwise, it's impossible to complete the circuit.

// 2. Greedy Strategy:

// Start at index 0 and track remaining fuel (tank).
// If at any point tank becomes negative, restart at the next station.
// The key observation is that if a valid solution exists, the station after a failure must be the new starting station.

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost, tank := 0, 0, 0
	startIndex := 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		tank += gas[i] - cost[i]

		if tank < 0 {
			// If tank is negative, reset startIndex and tank.
			startIndex = i + 1
			tank = 0
		}
	}

	if totalGas < totalCost {
		return -1
	}

	return startIndex
}
