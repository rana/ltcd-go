package main

import (
	"math"
	"testing"
)

// Time complexity: O(n), n is the length of the array prcs. We traverse the array once.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/67147424-275c-8002-a970-61aaec808a22
func maxProfit121(prcs []int) int {
	// Best Time to Buy and Sell Stock
	// Given integer array prcs.
	// Determine max profit.
	// Return max profit; or, zero.
	// Conditions:
	// * Buy once, sell once.
	// Use local optimization "greedy" approach.

	// Initialize variable.
	min_prc := math.MaxInt32
	max_pft := 0

	// Iterate through each element of array prcs.
	for _, prc := range prcs {
		// Check for minium price.
		if prc < min_prc {
			min_prc = prc
		} else {
			// Calculate profit.
			pft := prc - min_prc
			// Check for maximum profit.
			if pft > max_pft {
				max_pft = pft
			}
		}
	}

	return max_pft
}

func TestMaxProfit121(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Valid case - profit possible",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "Valid case - decreasing prices",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "Valid case - increasing prices",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Edge case - single price",
			prices:   []int{5},
			expected: 0,
		},
		{
			name:     "Edge case - all same prices",
			prices:   []int{0, 0, 0, 0},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit121(tt.prices)
			if result != tt.expected {
				t.Errorf("maxProfit1() = %d, want %d", result, tt.expected)
			}
		})
	}
}
