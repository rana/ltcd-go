package main

import (
	"testing"
)

// Time complexity: O(n), n is the length of the arrays. We traverse each array once.
// Space complexity: O(1), constant additional space used.
// https://console.anthropic.com/workbench/fc2821b1-dab9-4f65-8725-0911cd10b2d3
func canCompleteCircuit134(gas []int, cost []int) int {
	if len(gas) == 0 || len(gas) != len(cost) {
		return -1
	}

	totGas, curGas := 0, 0
	startIndex := 0

	for idx := 0; idx < len(gas); idx++ {
		dif := gas[idx] - cost[idx]
		totGas += dif
		curGas += dif

		if curGas < 0 {
			curGas = 0
			startIndex = idx + 1
		}
	}

	if totGas >= 0 {
		return startIndex
	}
	return -1
}

func TestCanCompleteCircuit134(t *testing.T) {
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
			result := canCompleteCircuit134(tt.gas, tt.cost)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
