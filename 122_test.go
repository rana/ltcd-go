package main

import "testing"

// Time complexity: O(n), n is the length of the array prcs. We traverse the array once.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/6714797e-7598-8002-b8ea-f625fcd45f64
func maxProfit122(prcs []int) int {
	// Best Time to Buy and Sell Stock II
	// Given an integer array prcs.
	// Determine max profit.
	// Return max profit.
	// Conditions:
	// * Hold one share at a time
	// * Buy one day, sell next day
	// * Repeat
	// Use local optimization "greedy" approach.

	// Initialize variable.
	maxPrf := 0

	// Iterate through each element of the prcs array.
	for idx := 1; idx < len(prcs); idx++ {
		// Check whether there is a daily profit.
		if prcs[idx] > prcs[idx-1] {
			// Accumulate daily profit.
			maxPrf += prcs[idx] - prcs[idx-1]
		}
	}

	return maxPrf
}

func TestMaxProfit122(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Ascending prices",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Descending prices",
			prices:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "Mixed prices",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			name:     "Single day",
			prices:   []int{5},
			expected: 0,
		},
		{
			name:     "Empty prices",
			prices:   []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit122(tt.prices)
			if got != tt.expected {
				t.Errorf("maxProfit122() = %d, want %d", got, tt.expected)
			}
		})
	}
}
