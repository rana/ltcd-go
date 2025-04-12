package main

import (
	"testing"
)

// Split Array Equal Sum
//
// Given a positive integer array nums.
// Determine whether it's possible to split
// into two subarrays with equal sum values.
// Return split idx; or, -1.
func splitArrayEqualSum1b(nums []int) int {
	return 0
}

func TestSplitArrayEqualSum1b(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 2, 3, 4, 5, 5},
			expected: 3,
		},
		{
			name:     "Example 2",
			input:    []int{4, 1, 2, 3},
			expected: 1,
		},
		{
			name:     "Example 3",
			input:    []int{4, 3, 2, 1},
			expected: -1,
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: -1,
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: -1,
		},
		{
			name:     "Two elements with equal sum",
			input:    []int{5, 5},
			expected: 0,
		},
		{
			name:     "All zeros",
			input:    []int{0, 0, 0, 0},
			expected: 0, // Any index works
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := splitArrayEqualSum1b(tc.input)
			if got != tc.expected {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
