package main

import (
	"testing"
)

// Time complexity:
// Space complexity:
func continuousSubarrays2762b(nums []int) int64 {
	return 0
}

func TestContinuousSubarrays2762b(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		// Standard examples from problem statement
		{
			name:     "Example 1: [5,4,2,4]",
			nums:     []int{5, 4, 2, 4},
			expected: 8,
		},
		{
			name:     "Example 2: [1,2,3]",
			nums:     []int{1, 2, 3},
			expected: 6,
		},

		// Edge cases
		{
			name:     "Single element",
			nums:     []int{7},
			expected: 1,
		},
		{
			name:     "Two identical elements",
			nums:     []int{3, 3},
			expected: 3, // [3], [3], [3,3]
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},

		// Array with all identical elements
		{
			name:     "All same elements",
			nums:     []int{5, 5, 5, 5},
			expected: 10, // n*(n+1)/2 = 4*5/2 = 10
		},

		// Special cases with max difference = 2
		{
			name:     "Max difference exactly 2",
			nums:     []int{10, 12, 10, 12, 10},
			expected: 15, // All possible subarrays are valid
		},
		{
			name:     "Three distinct values with diff ≤ 2",
			nums:     []int{1, 2, 3, 2, 1},
			expected: 15, // All possible subarrays are valid
		},

		// Cases where window needs to shrink
		{
			name:     "Requires window shrinking",
			nums:     []int{1, 5, 4, 2, 4},
			expected: 9, // First two elements form discontinuous subarray
		},
		{
			name:     "Alternating pattern forcing small windows",
			nums:     []int{1, 5, 1, 5, 1},
			expected: 5, // Only single-element subarrays are valid
		},

		// Increasing/decreasing sequences
		{
			name:     "Increasing sequence",
			nums:     []int{1, 3, 5, 7, 9},
			expected: 9, // Only windows of size 1 and 2 are valid
		},
		{
			name:     "Decreasing sequence",
			nums:     []int{9, 7, 5, 3, 1},
			expected: 9, // Only windows of size 1 and 2 are valid
		},

		// Repeated patterns
		{
			name:     "Repeated pattern within range",
			nums:     []int{1, 2, 3, 1, 2, 3},
			expected: 21, // All subarrays are valid
		},

		// Large value differences
		{
			name:     "Large jumps between values",
			nums:     []int{100, 1000, 10000, 100000},
			expected: 4, // Only individual elements form valid subarrays
		},

		// Special patterns
		{
			name:     "Plateau pattern",
			nums:     []int{1, 3, 3, 3, 1},
			expected: 15, // All possible subarrays are valid
		},
		{
			name:     "Valley pattern",
			nums:     []int{5, 3, 3, 3, 5},
			expected: 15, // All possible subarrays are valid
		},

		// Near boundary cases
		{
			name:     "At boundary",
			nums:     []int{1, 3, 5, 7},
			expected: 7, // Only pairs and individual elements
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := continuousSubarrays2762b(tc.nums)
			if got != tc.expected {
				t.Errorf("continuousSubarrays(%v) = %v, want %v", tc.nums, got, tc.expected)
			}
		})
	}
}
