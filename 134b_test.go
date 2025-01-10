package main

import (
	"testing"
)

func canCompleteCircuit134b(gas []int, cost []int) int {
	// Gas Station
	// Given two integer arrays gas and cost.
	// Determine the starting gas station index.
	// Return the starting gas station index.
	// Conditions:
	// * Arrays are circular
	// * gas[i] is tank of gas
	// * cost[i] is cost to travel from i to i+1
	// Track the difference of gas - cost for current station and total trip.

	totGas, curGas := 0, 0
	startIdx := 0
	for idx := 0; idx < len(gas); idx++ {
		// Calculate difference.
		dif := gas[idx] - cost[idx]
		totGas += dif
		curGas += dif

		// Check starting gas station condition.
		if curGas < 0 {
			curGas = 0
			startIdx = idx + 1
		}
	}

	// Check success condition.
	if totGas >= 0 {
		return startIdx
	}
	return -1
}

func TestCanCompleteCircuit134b(t *testing.T) {
	tests := []struct {
		name     string
		gas      []int
		cost     []int
		expected int
	}{
		{
			name:     "Success case - start from station 3",
			gas:      []int{1, 2, 3, 4, 5},
			cost:     []int{3, 4, 5, 1, 2},
			expected: 3,
		},
		{
			name:     "Failure case - no solution",
			gas:      []int{2, 3, 4},
			cost:     []int{3, 4, 3},
			expected: -1,
		},
		{
			name:     "Single station case",
			gas:      []int{5},
			cost:     []int{5},
			expected: 0,
		},
		{
			name:     "Zero remaining gas case",
			gas:      []int{2, 3, 4},
			cost:     []int{2, 3, 4},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canCompleteCircuit134b(tt.gas, tt.cost)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
