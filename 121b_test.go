package main

import (
	"math"
	"testing"
)

// Time complexity: O(n), n is the length of the prcs array.
// Space complexity: O(1), constant additional space used.
func maxProfit121b(prcs []int) int {
	// Best Time to Buy and Sell Stock
	// Given an integer array prcs.
	// Determine the maximum profit.
	// Return the maximum profit.
	// Conditions:
	// * Buy once and sell once.
	// Use a local optimization "greedy" strategy.

	// Initialize variables.
	min_prc := math.MaxInt64
	max_prf := 0

	// Iterate through prices.
	for _, prc := range prcs {
		// Check for min price.
		if prc < min_prc {
			min_prc = prc
		} else {
			// Calculate current profit.
			prf := prc - min_prc

			// Check max profit.
			if prf > max_prf {
				max_prf = prf
			}
		}
	}

	return max_prf
}

func TestMaxProfit121b(t *testing.T) {
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
			result := maxProfit121b(tt.prices)
			if result != tt.expected {
				t.Errorf("maxProfit1() = %d, want %d", result, tt.expected)
			}
		})
	}
}
