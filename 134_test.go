package main

import (
	"testing"
)

// Time complexity: O(n), n is the length of the arrays. We traverse each array once.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/67158357-55d0-8002-82a4-b231ea5d16c1
func canCompleteCircuit(gas []int, cost []int) int {
	// Gas Station
	// Given two integer arrays gas and cost.
	// Determine the starting index of a gas station
	// which allows total circuit traversal.
	// Return the starting index; or, -1.
	// Key insight, if total gas is great or equal to total cost
	// a solution is possible.

	// Initialize variables.
	tot_tnk, cur_tnk := 0, 0
	str_idx := 0

	// Iterate through each element of the arrays.
	for idx := 0; idx < len(gas); idx++ {
		// Calculate the in-out balance.
		dif := gas[idx] - cost[idx]
		tot_tnk += dif
		cur_tnk += dif
		// Evaluate starting station index.
		if cur_tnk < 0 {
			// Reset starting station to next station.
			cur_tnk = 0
			str_idx = idx + 1
		}
	}

	// Evaluate total circuit.
	if tot_tnk >= 0 {
		// Solution possible.
		return str_idx
	}
	return -1
}

func TestCanCompleteCircuit_Success(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	expected := 3

	result := canCompleteCircuit(gas, cost)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestCanCompleteCircuit_Failure(t *testing.T) {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	expected := -1

	result := canCompleteCircuit(gas, cost)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestCanCompleteCircuit_SingleStation(t *testing.T) {
	gas := []int{5}
	cost := []int{5}
	expected := 0

	result := canCompleteCircuit(gas, cost)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
