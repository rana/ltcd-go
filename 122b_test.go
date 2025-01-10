package main

import "testing"

func maxProfit122b(prcs []int) int {
	// Best Time to Buy and Sell Stock II
	// Given an integer array prcs.
	// Determine the maximum profit.
	// Return the maximum profit.
	// Conditions:
	// * Hold one share at a time.
	// * Can buy then sell multiple times.
	// Use a local optimization "greedy" algorithm

	maxPrf := 0
	for idx := 1; idx < len(prcs); idx++ {
		if prcs[idx] > prcs[idx-1] {
			maxPrf += prcs[idx] - prcs[idx-1]
		}
	}
	return maxPrf
}

func TestMaxProfit122b(t *testing.T) {
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
			got := maxProfit122b(tt.prices)
			if got != tt.expected {
				t.Errorf("maxProfit122() = %d, want %d", got, tt.expected)
			}
		})
	}
}
